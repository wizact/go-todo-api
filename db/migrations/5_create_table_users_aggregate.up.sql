CREATE TABLE IF NOT EXISTS users_aggregate (
    user_id TEXT NOT NULL PRIMARY KEY,
    value_data TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    deleted_at TEXT
) WITHOUT ROWID;