package mddb

import (
	"database/sql"

	"github.com/byliuyang/app/fw"
	migrate "github.com/rubenv/sql-migrate"
)

var _ fw.DBMigrationTool = (*PostgresMigrationTool)(nil)

type PostgresMigrationTool struct {
}

func (p PostgresMigrationTool) Migrate(db *sql.DB, migrationRoot string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationRoot,
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return err
}

func NewPostgresMigrationTool() PostgresMigrationTool {
	return PostgresMigrationTool{}
}
