CREATE TABLE IF NOT EXISTS users_token_view (
    user_id TEXT NOT NULL PRIMARY KEY,
    verification_token TEXT NOT NULL,
    verification_salt TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    deleted_at TEXT
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS verification_token_asc ON users_token_view (
    verification_token ASC);
