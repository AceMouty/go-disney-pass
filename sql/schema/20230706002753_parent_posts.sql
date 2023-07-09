-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS park_areas(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS area_rides(
  id SERIAL PRIMARY KEY,
  area_id INTEGER REFERENCES park_areas(id) ON DELETE CASCADE,
  ride_name VARCHAR(255) NOT NULL,
  CONSTRAINT fk_area_rides_area_id FOREIGN KEY (area_id) REFERENCES park_areas(id)
);

CREATE TABLE IF NOT EXISTS parent_posts(
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  parent_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  area_id INTEGER REFERENCES park_areas(id) ON DELETE CASCADE,
  ride_id INTEGER REFERENCES area_rides(id) ON DELETE CASCADE,
  is_open BOOLEAN NOT NULL DEFAULT FALSE,
  ride_time TIMESTAMP NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  number_of_kids INTEGER NOT NULL,
  CONSTRAINT fk_parent_posts_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
  CONSTRAINT fk_parent_posts_parent_id FOREIGN KEY (parent_id) REFERENCES users(user_id),
  CONSTRAINT fk_parent_posts_area_id FOREIGN KEY (area_id) REFERENCES park_areas(id),
  CONSTRAINT fk_parent_posts_ride_id FOREIGN KEY (ride_id) REFERENCES area_rides(id)
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS parent_posts;
DROP TABLE IF EXISTS area_rides;
DROP TABLE IF EXISTS park_areas;