package scaffolding

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"testing/fstest"
)

func fixtureConfig(t *testing.T) Config {
	t.Helper()
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "internal"), 0755); err != nil {
		t.Fatal(err)
	}
	return Config{Root: root, ModuleName: "todo", ImportPath: "example.com/project", Templates: fstest.MapFS{
		"module.go.tmpl":     &fstest.MapFile{Data: []byte("package {{.ModuleName}}\nvar ImportPath = {{printf \"%q\" .ImportPath}}\n")},
		"domain/doc.go.tmpl": &fstest.MapFile{Data: []byte("package domain\n")},
	}}
}

func TestBuildPlanDeterministic(t *testing.T) {
	c := fixtureConfig(t)
	first, err := BuildPlan(c)
	if err != nil {
		t.Fatal(err)
	}
	second, err := BuildPlan(c)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(first, second) {
		t.Fatal("plans differ")
	}
	if first.Files[0].Path != "domain/doc.go" {
		t.Fatal("files are not sorted")
	}
	if _, err := os.Stat(filepath.Join(c.Root, "internal", "todo")); !os.IsNotExist(err) {
		t.Fatal("planning wrote a module")
	}
}

func TestBuildPlanRejectsInvalidInputs(t *testing.T) {
	for _, name := range []string{"", "../todo", "todo/api", "123todo", "Todo", "todo_item", "type", "main", "test"} {
		t.Run(name, func(t *testing.T) {
			c := fixtureConfig(t)
			c.ModuleName = name
			if _, err := BuildPlan(c); err == nil {
				t.Fatal("accepted invalid name")
			}
		})
	}
}

func TestBuildPlanRejectsBrokenTemplates(t *testing.T) {
	for _, source := range []string{"{{", "package {{.Missing}}", "invalid Go"} {
		c := fixtureConfig(t)
		c.Templates = fstest.MapFS{"module.go.tmpl": &fstest.MapFile{Data: []byte(source)}}
		if _, err := BuildPlan(c); err == nil {
			t.Fatalf("accepted %q", source)
		}
	}
	c := fixtureConfig(t)
	c.Templates = fstest.MapFS{}
	if _, err := BuildPlan(c); err == nil {
		t.Fatal("accepted empty templates")
	}
}

func TestBuildPlanRejectsExistingModule(t *testing.T) {
	c := fixtureConfig(t)
	if err := os.Mkdir(filepath.Join(c.Root, "internal", "todo"), 0755); err != nil {
		t.Fatal(err)
	}
	if _, err := BuildPlan(c); err == nil {
		t.Fatal("accepted existing module")
	}
}

func TestBuildPlanRejectsMissingContext(t *testing.T) {
	for _, field := range []string{"root", "import path", "templates"} {
		c := fixtureConfig(t)
		switch field {
		case "root":
			c.Root = ""
		case "import path":
			c.ImportPath = ""
		case "templates":
			c.Templates = nil
		}
		if _, err := BuildPlan(c); err == nil {
			t.Fatalf("accepted missing %s", field)
		}
	}
}

func TestBuildPlanRejectsSymlinkDestination(t *testing.T) {
	c := fixtureConfig(t)
	if err := os.Symlink(t.TempDir(), filepath.Join(c.Root, "internal", "todo")); err != nil {
		t.Fatal(err)
	}
	if _, err := BuildPlan(c); err == nil {
		t.Fatal("accepted existing symlink")
	}
}
