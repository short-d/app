package db

import "database/sql"

type Connector interface {
	Connect(config Config) (*sql.DB, error)
}
