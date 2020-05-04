package dbtest

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/short-d/app/fw/db"
)

type dbConsumer func(sqlDB *sql.DB)

func AccessTestDB(
	dbConnector db.Connector,
	dbMigrationTool db.MigrationTool,
	dbMigrationRoot string,
	dbConfig db.Config,
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

	err = resetDatabase(db, dbMigrationRoot, dbMigrationTool)
	if err != nil {
		panic(err)
	}

	err = dbMigrationTool.MigrateUp(db, dbMigrationRoot)
	if err != nil {
		panic(err)
	}

	consumer(db)
}

func resetDatabase(db *sql.DB, dbMigrationRoot string, dbMigrationTool db.MigrationTool) error {
	err := dbMigrationTool.MigrateUp(db, dbMigrationRoot)
	if err != nil {
		return err
	}

	return dbMigrationTool.MigrateDown(db, dbMigrationRoot)
}
