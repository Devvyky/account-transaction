package db

import (
	"database/sql"
)

// Store provides all functions to execute db queries
type Store interface {
	Querier
}

// SQLStore provides all functions to execute db queries
type SQLStore struct {
	db *sql.DB
	*Queries
}

// Creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
