package scaffolding

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWritePlanCreatesModuleAndRefusesOverwrite(t *testing.T) {
	c := fixtureConfig(t)
	plan, err := BuildPlan(c)
	if err != nil {
		t.Fatal(err)
	}
	if err := WritePlan(c, plan); err != nil {
		t.Fatal(err)
	}
	for _, file := range plan.Files {
		got, err := os.ReadFile(filepath.Join(c.Root, "internal", c.ModuleName, filepath.FromSlash(file.Path)))
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != string(file.Content) {
			t.Fatalf("wrong content for %s", file.Path)
		}
	}
	if err := WritePlan(c, plan); err == nil {
		t.Fatal("overwrote existing module")
	}
}

func TestWritePlanRejectsUnsafePaths(t *testing.T) {
	for _, p := range []string{"../escape", "/absolute", ".", "a/../b", "a\\b"} {
		c := fixtureConfig(t)
		if err := WritePlan(c, Plan{Files: []File{{Path: p}}}); err == nil {
			t.Fatalf("accepted %q", p)
		}
		if _, err := os.Stat(filepath.Join(c.Root, "internal", "todo")); !os.IsNotExist(err) {
			t.Fatal("invalid plan wrote module")
		}
	}
}

func TestWritePlanCleansUpPartialFailure(t *testing.T) {
	c := fixtureConfig(t)
	plan := Plan{Files: []File{{Path: "a", Content: []byte("first")}, {Path: "a/b", Content: []byte("second")}}}
	if err := WritePlan(c, plan); err == nil {
		t.Fatal("expected parent/file conflict")
	}
	if _, err := os.Stat(filepath.Join(c.Root, "internal", "todo")); !os.IsNotExist(err) {
		t.Fatal("partial module remains")
	}
}

func TestWritePlanRefusesEscapingInternalSymlink(t *testing.T) {
	c := fixtureConfig(t)
	outside := t.TempDir()
	if err := os.Remove(filepath.Join(c.Root, "internal")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(c.Root, "internal")); err != nil {
		t.Fatal(err)
	}
	if err := WritePlan(c, Plan{Files: []File{{Path: "module.go", Content: []byte("package todo\n")}}}); err == nil {
		t.Fatal("followed escaping symlink")
	}
	entries, err := os.ReadDir(outside)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Fatal("wrote outside repository")
	}
}

func TestWritePlanRejectsEmptyOrDuplicatePlan(t *testing.T) {
	c := fixtureConfig(t)
	for _, plan := range []Plan{{}, {Files: []File{{Path: "a"}, {Path: "a"}}}} {
		if err := WritePlan(c, plan); err == nil {
			t.Fatal("accepted invalid plan")
		}
	}
}

func TestWritePlanPreservesExistingFiles(t *testing.T) {
	c := fixtureConfig(t)
	destination := filepath.Join(c.Root, "internal", "todo")
	if err := os.Mkdir(destination, 0755); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(destination, "keep.txt")
	if err := os.WriteFile(marker, []byte("original"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := WritePlan(c, Plan{Files: []File{{Path: "keep.txt", Content: []byte("replacement")}}}); err == nil {
		t.Fatal("accepted existing module")
	}
	got, err := os.ReadFile(marker)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "original" {
		t.Fatal("changed existing content")
	}
}
