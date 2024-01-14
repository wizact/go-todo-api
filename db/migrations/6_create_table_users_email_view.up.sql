CREATE TABLE IF NOT EXISTS users_email_view (
    user_id TEXT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    has_verified_email INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    deleted_at TEXT
) WITHOUT ROWID;