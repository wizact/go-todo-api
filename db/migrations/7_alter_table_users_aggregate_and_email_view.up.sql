DROP TABLE IF EXISTS users_aggregate;
CREATE TABLE IF NOT EXISTS users_aggregate (
    user_id TEXT PRIMARY KEY,
    value_data TEXT NOT NULL,
    version_info NUMBER,
    created_at INTEGER NOT NULL,
    updated_at INTEGER,
    deleted_at  INTEGER
) WITHOUT ROWID;

DROP TABLE IF EXISTS users_email_view;
CREATE TABLE IF NOT EXISTS users_email_view (
    user_id TEXT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    has_verified_email INTEGER NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER,
    deleted_at INTEGER
) WITHOUT ROWID;