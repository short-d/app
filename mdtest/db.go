package mdtest

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/short-d/app/fw"
)

type dbConsumer func(sqlDB *sql.DB)

func AccessTestDB(
	dbConnector fw.DBConnector,
	dbMigrationTool fw.DBMigrationTool,
	dbMigrationRoot string,
	dbConfig fw.DBConfig,
	consumer dbConsumer) {

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dbMigrationRoot = filepath.Join(workDir, dbMigrationRoot)

	db, err := dbConnector.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = dbMigrationTool.MigrateUp(db, dbMigrationRoot)
	if err != nil {
		panic(err)
	}

	consumer(db)

	err = dbMigrationTool.MigrateDown(db, dbMigrationRoot)
	if err != nil {
		panic(err)
	}
}
