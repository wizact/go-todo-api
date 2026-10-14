package scaffolding

import (
	"embed"
	"io/fs"
)

//go:embed templates/module.bundle.tmpl
var bundledTemplates embed.FS

// BundledTemplates returns the repository's architecture-and-wiring scaffold.
func BundledTemplates() fs.FS {
	templates, err := fs.Sub(bundledTemplates, "templates")
	if err != nil {
		panic(err)
	} // The embedded directory is fixed at compile time.
	return templates
}
