// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: arearides.sql

package database

import (
	"context"
)

const getAllAreaRides = `-- name: GetAllAreaRides :many
SELECT id, area_id, ride_name FROM area_rides
`

func (q *Queries) GetAllAreaRides(ctx context.Context) ([]AreaRide, error) {
	rows, err := q.db.QueryContext(ctx, getAllAreaRides)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AreaRide
	for rows.Next() {
		var i AreaRide
		if err := rows.Scan(&i.ID, &i.AreaID, &i.RideName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}