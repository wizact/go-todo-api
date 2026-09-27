CREATE TABLE IF NOT EXISTS user_registration_verifications (
    user_id TEXT NOT NULL PRIMARY KEY,
    secret_digest TEXT NOT NULL,
    expires_at INTEGER NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
) WITHOUT ROWID;
