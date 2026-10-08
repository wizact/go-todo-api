# Naming conventions

Project-owned filenames and folders use lowercase letters and digits with no word delimiters. Go tests retain the required `_test.go` suffix. Package names match their folders; external test packages retain `_test`.

Go identifiers use MixedCaps or mixedCaps, with standard initialisms such as ID, DB, API, HTTP, TLS, NATS and SQLite. Types describe their role without repeating package context; simple accessors omit Get, and operations that retrieve external data use Fetch or Find. Receivers are short and consistent.

Keep tool-required names (README.md, CLAUDE.md, LICENSE, Makefile, Dockerfile, Go module files and dot configuration), migration version/direction filenames and generated code under their owning tool conventions. Preserve JSON fields, SQL identifiers, environment variables, HTTP routes, CLI flags and event subjects.

The development-container files retain their existing names because this session cannot edit that directory. Ignored local files are untouched. REST-client examples and local environment configuration now live in `httpclient/`.
