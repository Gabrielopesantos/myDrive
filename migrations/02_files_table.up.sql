CREATE TABLE files (
    file_id UUID PRIMARY KEY,
    file_owner_id UUID NOT NULL REFERENCES users (user_id),
    file VARCHAR(256) NOT NULL CHECK (file <> ''),
    filename VARCHAR(256) NOT NULL CHECK (filename <> ''),
    description VARCHAR(1024) DEFAULT '',
    extension VARCHAR(32) NOT NULL,
    size FLOAT NOT NULL,
    tags VARCHAR(256) DEFAULT '',

    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP
)
