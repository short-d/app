package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Connector interface {
	Connect(config Config) (*sql.DB, error)
}

type PostgresConnector struct {
}

func (p PostgresConnector) Connect(config Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DbName,
	)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresConnector() PostgresConnector {
	return PostgresConnector{}
}
