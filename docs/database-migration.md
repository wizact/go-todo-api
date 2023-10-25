# Database Migration

## Requirements:
This process requires go-bindata to be installed and accessible in the user path:

### Install go-bindata
```
go install github.com/go-bindata/go-bindata/go-bindata
```

### Step 1: Create migration files
In the `./db/migrations` create a two files for each change. The naming of the files should follow the following format:

`[sequence_number]_description_[up|down].sql`

`sequence_number` is an incremental integer value that is used for tracking the current state of our migration. After successful migration, the maximum sequence number is persisted in the database versioning table.

`up` indicates the scripts that will be called in the migration upgrade process. `down` scripts indicates the rollback process - if necessary.

### Step 2: Create resource file
We use `go-bindata` to embed the migration files in a go file. This enables us to build a standalone executable, deploy it to a environment, and execute the migration process remotely.

In order to create the resource files, execute the following command:

```
./db/migration.sh
```

### Step 3: Execute the migration
You can execute the migration, or create the binrary file using the `cmd/db-migration` entrypoint.

```
go run ./cmd/db-migration.go
```