package mddb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app/fw"
	_ "github.com/lib/pq"
)

var _ fw.DBConnector = (*PostgresConnector)(nil)

type PostgresConnector struct {
}

func (p PostgresConnector) Connect(dbConfig fw.DBConfig) (*sql.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DbName,
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
