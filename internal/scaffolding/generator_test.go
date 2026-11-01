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

func TestBuildPlanRendersNamedTemplateBundle(t *testing.T) {
	c := fixtureConfig(t)
	c.Templates = fstest.MapFS{"architecture.bundle.tmpl": &fstest.MapFile{Data: []byte(`{{define "module.go"}}package {{.ModuleName}}{{end}}{{define "domain/model/doc.go"}}package model{{end}}`)}}
	plan, err := BuildPlan(c)
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Files) != 2 || plan.Files[0].Path != "domain/model/doc.go" || plan.Files[1].Path != "module.go" {
		t.Fatalf("unexpected files: %v", plan.Files)
	}
}

func TestBuildPlanRejectsUnsafeOrDuplicateBundleOutputs(t *testing.T) {
	for _, source := range []string{`{{define "../escape.go"}}package escape{{end}}`, `{{define "."}}empty{{end}}`} {
		c := fixtureConfig(t)
		c.Templates = fstest.MapFS{"bad.bundle.tmpl": &fstest.MapFile{Data: []byte(source)}}
		if _, err := BuildPlan(c); err == nil {
			t.Fatal("accepted unsafe bundle output")
		}
	}
	c := fixtureConfig(t)
	c.Templates = fstest.MapFS{"a.bundle.tmpl": &fstest.MapFile{Data: []byte(`{{define "module.go"}}package todo{{end}}`)}, "module.go.tmpl": &fstest.MapFile{Data: []byte("package todo")}}
	if _, err := BuildPlan(c); err == nil {
		t.Fatal("accepted duplicate output")
	}
}

func TestBuildPlanRendersOutputPaths(t *testing.T) {
	c := fixtureConfig(t)
	c.Templates = fstest.MapFS{"publisher.bundle.tmpl": &fstest.MapFile{Data: []byte(`{{define "ports/event/{{.ModuleName}}eventpublisher.go"}}package event
 type {{.ModuleTypeName}}EventPublisher interface {}{{end}}`)}}
	plan, err := BuildPlan(c)
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Files) != 1 || plan.Files[0].Path != "ports/event/todoeventpublisher.go" {
		t.Fatalf("unexpected files %v", plan.Files)
	}
}

func TestBuildPlanValidatesRenderedPaths(t *testing.T) {
	for _, tc := range []struct{ name, source string }{
		{"traversal", `{{define "{{.ImportPath}}/file.go"}}package event{{end}}`},
		{"missing value", `{{define "{{.Missing}}/file.go"}}package event{{end}}`},
		{"duplicate", `{{define "{{.ModuleName}}.go"}}package todo{{end}}{{define "todo.go"}}package todo{{end}}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := fixtureConfig(t)
			c.ImportPath = "../outside"
			c.Templates = fstest.MapFS{"outputs.bundle.tmpl": &fstest.MapFile{Data: []byte(tc.source)}}
			if _, err := BuildPlan(c); err == nil {
				t.Fatal("accepted invalid rendered output path")
			}
		})
	}
}
