# Database migrations

SQLite schema changes are stored as ordered `up` and `down` scripts under `db/migrations`. The scripts are embedded in `db/resourcefile.go` and applied by the `cmd/dbmigration` executable through `golang-migrate`.

## Requirements

Use the Go 1.27 toolchain configured by the project. Resource generation requires `go-bindata`; `build/migration.sh` installs it when it is not available.

## Create a migration

Add two files using the existing naming convention:

```text
<sequence>_<description>.up.sql
<sequence>_<description>.down.sql
```

The `up` script applies one schema change. The `down` script reverses that change when reversal is safe. Sequence numbers must be monotonically increasing.

Migration scripts do not require test-first development. Review both directions, regenerate the embedded resource, and execute the migration against a disposable database before using it on persistent data.

## Generate the embedded resource

```sh
make gen-db-resource
```

This recreates `db/resourcefile.go` from every SQL file under `db/migrations`. Commit the generated resource whenever migration inputs change.

## Apply migrations

Apply all pending migrations to the repository's development database:

```sh
make run-db-migration
```

To select another database explicitly:

```sh
TODOAPI_DBPATH=/path/to/todo.db go run ./cmd/dbmigration
```

The migration command treats `no change` as a successful no-op.

## Build the migration executable

```sh
make build-dbmigration OS=linux ARCH=amd64
```

## Registration verification schema

Migration 10 introduced `user_registration_verifications`:

| Column | Purpose |
| --- | --- |
| `user_id` | Primary key and aggregate identity |
| `secret_digest` | Bcrypt digest; the raw secret is never stored |
| `expires_at` | Expiry as a Unix millisecond timestamp |
| `created_at` | Creation time as a Unix millisecond timestamp |
| `updated_at` | Replacement time as a Unix millisecond timestamp |

One row per user allows issuing a new credential to replace the previous credential.

Migration 11 dropped the legacy `users_token_view` and its token index. Legacy raw credentials are intentionally invalidated rather than migrated; affected users must receive a newly issued credential.

Runtime verification completion is separate from schema migration. `SQLiteUserRepository.CompleteRegistration` updates `users_aggregate.value_data`, updates `users_email_view.has_verified_email`, and deletes the matching verification row in one database transaction.
