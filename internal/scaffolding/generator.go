// Package scaffolding plans and creates new modules without changing existing code.
package scaffolding

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"go/token"
	"io/fs"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// Config supplies repository context and trusted scaffold templates.
type Config struct {
	Root       string
	ModuleName string
	ImportPath string
	Templates  fs.FS
}

// File is a rendered file, with a slash-separated path relative to the new module.
type File struct {
	Path    string
	Content []byte
}

// Plan contains sorted, fully rendered files. Building a plan performs no writes.
type Plan struct{ Files []File }

// TemplateData contains values available to both text and Go templates.
type TemplateData struct {
	ModuleName     string
	ModuleTypeName string
	ImportPath     string
}

var moduleNamePattern = regexp.MustCompile(`^[a-z][a-z0-9]*$`)

func validateConfig(c Config) error {
	if !moduleNamePattern.MatchString(c.ModuleName) || token.Lookup(c.ModuleName).IsKeyword() || c.ModuleName == "main" || c.ModuleName == "test" {
		return fmt.Errorf("invalid module name %q: use a lowercase Go package name without delimiters", c.ModuleName)
	}
	if c.Root == "" {
		return errors.New("repository root is required")
	}
	if c.ImportPath == "" || strings.ContainsAny(c.ImportPath, " \t\r\n\"\\") {
		return errors.New("valid repository import path is required")
	}
	return nil
}

func openInternal(c Config) (*os.Root, error) {
	root, err := os.OpenRoot(c.Root)
	if err != nil {
		return nil, fmt.Errorf("open repository: %w", err)
	}
	defer root.Close()
	internal, err := root.OpenRoot("internal")
	if err != nil {
		return nil, fmt.Errorf("open repository internal directory: %w", err)
	}
	return internal, nil
}

// BuildPlan renders .tmpl files and named outputs in .bundle.tmpl files,
// sorts output paths, and formats Go output.
func BuildPlan(c Config) (Plan, error) {
	if err := validateConfig(c); err != nil {
		return Plan{}, err
	}
	if c.Templates == nil {
		return Plan{}, errors.New("templates are required")
	}
	internal, err := openInternal(c)
	if err != nil {
		return Plan{}, err
	}
	defer internal.Close()
	if _, err := internal.Lstat(c.ModuleName); !errors.Is(err, fs.ErrNotExist) {
		if err != nil {
			return Plan{}, fmt.Errorf("check module destination: %w", err)
		}
		return Plan{}, fmt.Errorf("module %q already exists", c.ModuleName)
	}
	data := TemplateData{ModuleName: c.ModuleName, ModuleTypeName: strings.ToUpper(c.ModuleName[:1]) + c.ModuleName[1:], ImportPath: c.ImportPath}
	plan := Plan{}
	err = fs.WalkDir(c.Templates, ".", func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("template %q must be a regular file", name)
		}
		source, err := fs.ReadFile(c.Templates, name)
		if err != nil {
			return fmt.Errorf("read template %s: %w", name, err)
		}
		tmpl, err := template.New(name).Option("missingkey=error").Parse(string(source))
		if err != nil {
			return fmt.Errorf("parse template %s: %w", name, err)
		}
		outputs := map[string]string{strings.TrimSuffix(name, ".tmpl"): name}
		if strings.HasSuffix(name, ".bundle.tmpl") {
			outputs = make(map[string]string)
			for _, defined := range tmpl.Templates() {
				if defined.Name() != name {
					outputs[defined.Name()] = defined.Name()
				}
			}
		}
		destinations := make([]string, 0, len(outputs))
		for destination := range outputs {
			destinations = append(destinations, destination)
		}
		sort.Strings(destinations)
		for _, destination := range destinations {
			templateName := outputs[destination]
			pathTemplate, err := template.New("output path").Option("missingkey=error").Parse(destination)
			if err != nil {
				return fmt.Errorf("parse output path %q: %w", destination, err)
			}
			var renderedPath bytes.Buffer
			if err := pathTemplate.Execute(&renderedPath, data); err != nil {
				return fmt.Errorf("render output path %q: %w", destination, err)
			}
			destination = renderedPath.String()
			if destination == "." || !fs.ValidPath(destination) || strings.Contains(destination, "\\") || path.Clean(destination) != destination {
				return fmt.Errorf("invalid output path %q in template %s", destination, name)
			}
			var output bytes.Buffer
			if err := tmpl.ExecuteTemplate(&output, templateName, data); err != nil {
				return fmt.Errorf("render template %s: %w", templateName, err)
			}
			content := output.Bytes()
			if strings.HasSuffix(destination, ".go") {
				content, err = format.Source(content)
				if err != nil {
					return fmt.Errorf("format %s: %w", destination, err)
				}
			}
			plan.Files = append(plan.Files, File{destination, content})
		}
		return nil
	})
	if err != nil {
		return Plan{}, err
	}
	if len(plan.Files) == 0 {
		return Plan{}, errors.New("no .tmpl files found")
	}
	sort.Slice(plan.Files, func(i, j int) bool { return plan.Files[i].Path < plan.Files[j].Path })
	for i := 1; i < len(plan.Files); i++ {
		if plan.Files[i-1].Path == plan.Files[i].Path {
			return Plan{}, fmt.Errorf("duplicate output path %q", plan.Files[i].Path)
		}
	}
	return plan, nil
}
