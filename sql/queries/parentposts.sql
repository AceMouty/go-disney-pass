-- name: CreateParentPost :one
INSERT INTO parent_posts(
  user_id
  ,parent_id
  ,area_id
  ,ride_id
  ,is_open
  ,ride_time
  ,number_of_kids
)
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetParentPostById :one
SELECT 
pp.id
,u1.username as username
,u2.username as parentname
,pa.name as park_area_name
,ar.ride_name
,pp.is_open
,pp.ride_time
,pp.number_of_kids
FROM parent_posts as pp
JOIN users as u1 ON u1.user_id = pp.user_id
JOIN users as u2 ON u2.parent_id = pp.parent_id
JOIN park_areas as pa ON pa.id = pp.area_id
JOIN area_rides as ar ON ar.id = pp.ride_id
WHERE pp.id = $1;
