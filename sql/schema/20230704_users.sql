-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  user_id UUID PRIMARY KEY,
  username varchar(50) UNIQUE NOT NULL,
  password varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;
