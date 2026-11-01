package scaffolding

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestBundledTemplatesCreateCompilableArchitecture(t *testing.T) {
	c := fixtureConfig(t)
	c.Templates = BundledTemplates()
	plan, err := BuildPlan(c)
	if err != nil {
		t.Fatal(err)
	}
	paths := []string{}
	for _, file := range plan.Files {
		paths = append(paths, file.Path)
	}
	want := []string{"README.md", "adapters/controller/doc.go", "adapters/controller/model/doc.go", "adapters/http/doc.go", "adapters/repository/doc.go", "application/service/doc.go", "domain/aggregate/doc.go", "domain/model/doc.go", "domain/usecase/doc.go", "module.go", "ports/event/todoeventpublisher.go", "ports/repository/doc.go", "ports/service/doc.go", "ports/usecase/doc.go"}
	if !reflect.DeepEqual(paths, want) {
		t.Fatalf("architecture paths = %v", paths)
	}
	if err := WritePlan(c, plan); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(c.Root, "go.mod"), []byte("module example.com/project\n\ngo 1.27.0\n"), 0644); err != nil {
		t.Fatal(err)
	}
	// Exercise the generated constructor as a consumer, not just a syntax check.
	source := `package todo_test
import (
 "testing"
 todo "example.com/project/internal/todo"
 event "example.com/project/internal/todo/ports/event"
)
var _ event.TodoEventPublisher = struct{}{}
func TestConstruction(t *testing.T) {
 module, err := todo.NewModuleWithDependencies(todo.Dependencies{})
 if err != nil || module == nil { t.Fatalf("module=%v error=%v", module, err) }
 if module, err := todo.NewModuleWithDependencies(todo.Dependencies{}, nil); err == nil || module != nil { t.Fatal("accepted nil option") }
}
`
	if err := os.WriteFile(filepath.Join(c.Root, "internal", "todo", "module_test.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	command := exec.Command("go", "test", "./internal/todo/...")
	command.Dir = c.Root
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("generated tests: %v\n%s", err, output)
	}
}
