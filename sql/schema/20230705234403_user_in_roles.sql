-- +goose Up
CREATE TABLE IF NOT EXISTS user_in_roles (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user_in_roles_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_user_in_roles_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
);


-- +goose Down
DROP TABLE IF EXISTS user_in_roles;
