-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS roles;
