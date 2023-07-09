-- +goose Up
-- +goose StatementBegin
INSERT INTO park_areas(name)
VALUES
  ('Fantasy Land'),
  ('Tomorrow Land');

INSERT INTO area_rides(area_id, ride_name)
VALUES
  (1, 'Prince Charming Redal Carrousel'),
  (1, 'Mad Tea Party'),
  (2, 'Tomorrowland Speedway'),
  (2, 'Space Mountain');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE park_areas CASCADE;
ALTER SEQUENCE park_areas_id_seq RESTART;
ALTER SEQUENCE area_rides_id_seq RESTART;
-- +goose StatementEnd
