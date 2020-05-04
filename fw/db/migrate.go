package db

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

type MigrationTool interface {
	MigrateUp(db *sql.DB, migrationRoot string) error
	MigrateDown(db *sql.DB, migrationRoot string) error
}

var _ MigrationTool = (*PostgresMigrationTool)(nil)

type PostgresMigrationTool struct {
}

func (p PostgresMigrationTool) MigrateUp(db *sql.DB, migrationRoot string) error {
	return p.migrate(db, migrationRoot, migrate.Up)
}

func (p PostgresMigrationTool) MigrateDown(db *sql.DB, migrationRoot string) error {
	return p.migrate(db, migrationRoot, migrate.Down)
}

func (p PostgresMigrationTool) migrate(
	db *sql.DB,
	migrationRoot string,
	migrateDirection migrate.MigrationDirection,
) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationRoot,
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrateDirection)
	return err
}

func NewPostgresMigrationTool() PostgresMigrationTool {
	return PostgresMigrationTool{}
}
