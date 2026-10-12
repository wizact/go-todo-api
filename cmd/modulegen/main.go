// Command modulegen previews or creates a module from trusted templates.
package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/wizact/go-todo-api/internal/scaffolding"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil && !errors.Is(err, flag.ErrHelp) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("modulegen", flag.ContinueOnError)
	flags.SetOutput(stderr)
	name := flags.String("name", "", "lowercase module name")
	root := flags.String("root", ".", "repository root containing go.mod and internal/")
	templates := flags.String("templates", "", "trusted template directory (required until bundled templates are added)")
	dryRun := flags.Bool("dry-run", false, "preview paths without writing files")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("unexpected positional arguments")
	}
	if *name == "" || *templates == "" {
		return errors.New("-name and -templates are required")
	}
	// Let the Go tool resolve the module directive instead of parsing go.mod here.
	command := exec.Command("go", "list", "-m", "-f", "{{.Path}}")
	command.Dir = *root
	importPath, err := command.Output()
	if err != nil {
		return fmt.Errorf("resolve repository import path: %w", err)
	}
	templateRoot, err := os.OpenRoot(*templates)
	if err != nil {
		return fmt.Errorf("open templates: %w", err)
	}
	defer templateRoot.Close()
	config := scaffolding.Config{Root: *root, ModuleName: *name, ImportPath: strings.TrimSpace(string(importPath)), Templates: templateRoot.FS()}
	plan, err := scaffolding.BuildPlan(config)
	if err != nil {
		return err
	}
	if !*dryRun {
		if err := scaffolding.WritePlan(config, plan); err != nil {
			return err
		}
	}
	action := "created"
	if *dryRun {
		action = "would create"
	}
	for _, file := range plan.Files {
		if _, err := fmt.Fprintf(stdout, "%s internal/%s/%s\n", action, *name, file.Path); err != nil {
			return err
		}
	}
	return nil
}
