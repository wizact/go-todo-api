package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func cliFixture(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	templates := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module example.com/project\n\ngo 1.27.0\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "internal"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(templates, "module.go.tmpl"), []byte("package {{.ModuleName}}\nvar ImportPath = {{printf \"%q\" .ImportPath}}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	return root, templates
}

func TestRunDryRunDoesNotWrite(t *testing.T) {
	root, templates := cliFixture(t)
	var output bytes.Buffer
	if err := run([]string{"-root", root, "-templates", templates, "-name", "todo", "-dry-run"}, &output, &output); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "internal/todo/module.go") {
		t.Fatal("missing preview")
	}
	if _, err := os.Stat(filepath.Join(root, "internal", "todo")); !os.IsNotExist(err) {
		t.Fatal("dry run wrote module")
	}
}

func TestRunCreatesFormattedModule(t *testing.T) {
	root, templates := cliFixture(t)
	var output bytes.Buffer
	if err := run([]string{"-root", root, "-templates", templates, "-name", "todo"}, &output, &output); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(root, "internal", "todo", "module.go"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "package todo\n\nvar ImportPath = \"example.com/project\"\n" {
		t.Fatalf("unexpected output %q", got)
	}
	if err := run([]string{"-root", root, "-templates", templates, "-name", "todo"}, &output, &output); err == nil {
		t.Fatal("repeated generation accepted")
	}
}

func TestRunRejectsInvalidArguments(t *testing.T) {
	for _, args := range [][]string{{}, {"-unknown"}, {"-name", "todo", "unexpected"}} {
		var output bytes.Buffer
		if err := run(args, &output, &output); err == nil {
			t.Fatalf("accepted %v", args)
		}
	}
}

func TestRunUsesBundledTemplatesByDefault(t *testing.T) {
	root, _ := cliFixture(t)
	var output bytes.Buffer
	if err := run([]string{"-root", root, "-name", "todo", "-dry-run"}, &output, &output); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"module.go", "domain/model/doc.go", "adapters/http/doc.go", "ports/output/repository/doc.go"} {
		if !strings.Contains(output.String(), "internal/todo/"+name) {
			t.Fatalf("missing %s", name)
		}
	}
	if _, err := os.Stat(filepath.Join(root, "internal", "todo")); !os.IsNotExist(err) {
		t.Fatal("dry run wrote module")
	}
}
