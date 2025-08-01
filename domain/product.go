package domain

import "database/sql"

type Product struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Price     string       `db:"price"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
