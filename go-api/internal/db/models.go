// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Authors struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}
