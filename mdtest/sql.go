package mdtest

import (
	"database/sql"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
)

type SQLStub struct {
	mock sqlmock.Sqlmock
}

func (s SQLStub) ExpectQuery(expectedSQL string) *ExpectedQuery {
	return &ExpectedQuery{query: s.mock.ExpectQuery(expectedSQL)}
}

func (s SQLStub) ExpectExec(expectedSQL string) *ExpectedExec {
	return &ExpectedExec{exec: s.mock.ExpectExec(expectedSQL)}
}

func NewSQLStub() (*sql.DB, SQLStub, error) {
	db, mock, err := sqlmock.New()
	return db, SQLStub{mock: mock}, err
}

type ExpectedQuery struct {
	query *sqlmock.ExpectedQuery
}

func (e ExpectedQuery) WillReturnRows(rows *TableRows) *ExpectedQuery {
	return &ExpectedQuery{query: e.query.WillReturnRows(rows.mockRows)}
}

type ExpectedExec struct {
	exec *sqlmock.ExpectedExec
}

func (e ExpectedExec) WillReturnError(err error) *ExpectedExec {
	return &ExpectedExec{exec: e.exec.WillReturnError(err)}
}

func (e ExpectedExec) WillReturnResult(result driver.Result) *ExpectedExec {
	return &ExpectedExec{exec: e.exec.WillReturnResult(result)}
}

type TableRows struct {
	mockRows *sqlmock.Rows
}

func (t TableRows) AddRow(args ...driver.Value) *TableRows {
	return &TableRows{mockRows: t.mockRows.AddRow(args...)}
}

func NewTableRows(columns []string) *TableRows {
	return &TableRows{mockRows: sqlmock.NewRows(columns)}
}
