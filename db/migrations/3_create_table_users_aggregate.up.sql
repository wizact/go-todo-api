CREATE TABLE IF NOT EXISTS users_aggregate (
    aggregate_id TEXT PRIMARY KEY,
    value_data TEXT NOT NULL,
    version_info NUMBER,
    create_date TEXT NOT NULL,
    update_date TEXT
) WITHOUT ROWID;