package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeAcceptanceFile(t *testing.T, root, name string, content []byte) {
	t.Helper()
	destination := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(destination, content, 0644); err != nil {
		t.Fatal(err)
	}
}

func acceptanceRepository(t *testing.T, importPath string) string {
	t.Helper()
	root := t.TempDir()
	for name, content := range map[string]string{
		"go.mod":                              fmt.Sprintf("module %s\n\ngo 1.27.0\n", importPath),
		"internal/api/bootstrap.go":           "package api\n// Existing application bootstrap.\n",
		"internal/user/module.go":             "package user\n// Existing user implementation.\n",
		"db/migrations/000001_example.up.sql": "CREATE TABLE example (id TEXT);\n",
	} {
		writeAcceptanceFile(t, root, name, []byte(content))
	}
	return root
}

func snapshotRepository(t *testing.T, root string) map[string][]byte {
	t.Helper()
	files := make(map[string][]byte)
	if err := filepath.WalkDir(root, func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(root, name)
		if err != nil {
			return err
		}
		content, err := os.ReadFile(name)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(relative)] = content
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return files
}

func acceptanceCommand(t *testing.T, root, program string, args ...string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	command := exec.CommandContext(ctx, program, args...)
	command.Dir = root
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("%s %v: %v\n%s", program, args, err, output)
	}
	return string(output)
}

func acceptanceFailure(t *testing.T, root, want, program string, args ...string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	command := exec.CommandContext(ctx, program, args...)
	command.Dir = root
	command.Env = append(os.Environ(), "GOWORK=off")
	output, err := command.CombinedOutput()
	if err == nil || !strings.Contains(string(output), want) {
		t.Fatalf("%s %v: error=%v output=%s; want failure containing %q", program, args, err, output, want)
	}
}

func buildAcceptanceCLI(t *testing.T) string {
	t.Helper()
	binary := filepath.Join(t.TempDir(), "modulegen")
	acceptanceCommand(t, ".", "go", "build", "-o", binary, ".")
	return binary
}

func copyAcceptanceGenerator(t *testing.T, root string) {
	t.Helper()
	// Explicit allowlist copies only generator code, its embed, and entry points.
	// No repository walk or environment-file access is needed.
	for _, name := range []string{
		"cmd/modulegen/main.go", "internal/scaffolding/generator.go", "internal/scaffolding/write.go",
		"internal/scaffolding/templates.go", "internal/scaffolding/templates/module.bundle.tmpl",
		"generate.go", "Makefile",
	} {
		content, err := os.ReadFile(filepath.Join("..", "..", filepath.FromSlash(name)))
		if err != nil {
			t.Fatal(err)
		}
		writeAcceptanceFile(t, root, name, content)
	}
}
