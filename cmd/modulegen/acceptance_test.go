package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestAcceptanceCLIWorkflow(t *testing.T) {
	binary := buildAcceptanceCLI(t)
	for _, tc := range []struct{ name, importPath, publisherType string }{
		{"todo", "example.com/acme/project", "TodoEventPublisher"},
		{"inventory2", "example.org/team/service/v2", "Inventory2EventPublisher"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := acceptanceRepository(t, tc.importPath)
			original := snapshotRepository(t, root)
			preview := acceptanceCommand(t, root, binary, "-name", tc.name, "-dry-run")
			if next := acceptanceCommand(t, root, binary, "-name", tc.name, "-dry-run"); preview != next {
				t.Fatal("preview is not deterministic")
			}
			if !reflect.DeepEqual(original, snapshotRepository(t, root)) {
				t.Fatal("preview modified repository")
			}
			created := acceptanceCommand(t, root, binary, "-name", tc.name)
			if created != strings.ReplaceAll(preview, "would create ", "created ") {
				t.Fatal("created files differ from preview")
			}
			generatedRoot := filepath.Join(root, "internal", tc.name)
			generated := snapshotRepository(t, generatedRoot)
			if len(generated) != 14 {
				t.Fatalf("generated %d files, want 14", len(generated))
			}
			for name, content := range generated {
				if strings.HasSuffix(name, ".go") {
					formatted, err := format.Source(content)
					if err != nil {
						t.Fatal(err)
					}
					if !bytes.Equal(formatted, content) {
						t.Fatalf("unformatted Go file %s", name)
					}
				}
			}
			if !strings.Contains(string(generated["README.md"]), tc.importPath+"/internal/"+tc.name) {
				t.Fatal("README has wrong module import path")
			}
			publisher := string(generated["ports/event/"+tc.name+"eventpublisher.go"])
			if !strings.Contains(publisher, "type "+tc.publisherType+" interface {") {
				t.Fatalf("missing expected event publisher interface %s", tc.publisherType)
			}
			if !strings.Contains(publisher, "// PublishEvent(ctx context.Context, event domain.ExampleDomainEvent) error") {
				t.Fatal("missing commented event publication example")
			}
			acceptanceCommand(t, root, "go", "test", "./internal/"+tc.name+"/...")
			writeAcceptanceFile(t, root, "internal/"+tc.name+"/module.go", append(generated["module.go"], []byte("\n// Developer-owned edit.\n")...))
			edited := snapshotRepository(t, root)
			acceptanceFailure(t, root, "already exists", binary, "-name", tc.name)
			if !reflect.DeepEqual(edited, snapshotRepository(t, root)) {
				t.Fatal("repeat generation modified edited module")
			}
			// Compare all pre-existing files and reject any unexpected write outside the new module.
			for name, content := range snapshotRepository(t, root) {
				if strings.HasPrefix(name, "internal/"+tc.name+"/") {
					continue
				}
				if !bytes.Equal(original[name], content) {
					t.Fatalf("changed or added unrelated file %s", name)
				}
			}
			for name, content := range original {
				got, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
				if err != nil || !bytes.Equal(content, got) {
					t.Fatalf("changed or removed original file %s", name)
				}
			}
		})
	}
}

func TestAcceptanceGenerationEntryPoints(t *testing.T) {
	t.Setenv("MODULE_NAME", "scaffoldcheck")
	for _, entry := range []string{"go generate", "make"} {
		t.Run(entry, func(t *testing.T) {
			root := acceptanceRepository(t, "github.com/wizact/go-todo-api")
			copyAcceptanceGenerator(t, root)
			original := snapshotRepository(t, root)
			if entry == "make" {
				acceptanceCommand(t, root, "make", "preview-module", "MODULE_NAME=scaffoldcheck", "VERSION=acceptance")
				if !reflect.DeepEqual(original, snapshotRepository(t, root)) {
					t.Fatal("Makefile preview wrote files")
				}
				acceptanceCommand(t, root, "make", "gen-module", "MODULE_NAME=scaffoldcheck", "VERSION=acceptance")
			} else {
				acceptanceCommand(t, root, "go", "generate", "-run", "modulegen", ".")
			}
			acceptanceCommand(t, root, "go", "test", "./internal/scaffoldcheck/...")
			after := snapshotRepository(t, root)
			if entry == "make" {
				acceptanceFailure(t, root, "already exists", "make", "gen-module", "MODULE_NAME=scaffoldcheck", "VERSION=acceptance")
			} else {
				acceptanceFailure(t, root, "already exists", "go", "generate", "-run", "modulegen", ".")
			}
			if !reflect.DeepEqual(after, snapshotRepository(t, root)) {
				t.Fatal("repeat entry point changed repository")
			}
		})
	}
}
