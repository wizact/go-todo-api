package scaffolding

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

// WritePlan exclusively creates a module. Failure removes the directory created
// by this invocation; existing modules are never modified or removed.
func WritePlan(c Config, plan Plan) (err error) {
	if err := validateConfig(c); err != nil {
		return err
	}
	if len(plan.Files) == 0 {
		return errors.New("cannot write an empty plan")
	}
	seen := make(map[string]bool)
	for _, file := range plan.Files {
		if file.Path == "." || !fs.ValidPath(file.Path) || strings.Contains(file.Path, "\\") || seen[file.Path] {
			return fmt.Errorf("invalid or duplicate output path %q", file.Path)
		}
		seen[file.Path] = true
	}
	internal, err := openInternal(c)
	if err != nil {
		return err
	}
	defer internal.Close()
	if err := internal.Mkdir(c.ModuleName, 0755); err != nil {
		return fmt.Errorf("create module %s exclusively: %w", c.ModuleName, err)
	}
	defer func() {
		if err != nil {
			if cleanupErr := internal.RemoveAll(c.ModuleName); cleanupErr != nil {
				err = errors.Join(err, fmt.Errorf("remove partial module: %w", cleanupErr))
			}
		}
	}()
	module, err := internal.OpenRoot(c.ModuleName)
	if err != nil {
		return err
	}
	defer module.Close()
	for _, file := range plan.Files {
		if err := module.MkdirAll(path.Dir(file.Path), 0755); err != nil {
			return fmt.Errorf("create parent for %s: %w", file.Path, err)
		}
		output, err := module.OpenFile(file.Path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			return fmt.Errorf("create %s: %w", file.Path, err)
		}
		_, writeErr := output.Write(file.Content)
		closeErr := output.Close()
		if err := errors.Join(writeErr, closeErr); err != nil {
			return fmt.Errorf("write %s: %w", file.Path, err)
		}
	}
	return nil
}
