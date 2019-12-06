package mdtest

import (
	"database/sql"
	"github.com/byliuyang/app/fw"
	"os"
	"path/filepath"
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
	defer db.Close()
	if err != nil {
		panic(err)
	}

	err = resetDatabase(db, dbMigrationRoot, dbMigrationTool)

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

func resetDatabase(db *sql.DB, dbMigrationRoot string, dbMigrationTool fw.DBMigrationTool) error {
	err := dbMigrationTool.MigrateUp(db, dbMigrationRoot)
	if err != nil {
		return err
	}

	return dbMigrationTool.MigrateDown(db, dbMigrationRoot)
}