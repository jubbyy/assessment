//go:build unittest

package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnectDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	_ = mock
	if err != nil {
		panic(err.Error())
	}
	db.Exec(CREATE_TABLE)

	GetStmt, _ = db.Prepare(SELECT_ID)
	DelStmt, _ = db.Prepare(DELETE_ID)
	PostStmt, _ = db.Prepare(INSERT)
	PutStmt, _ = db.Prepare(UPDATE_ID)
	ListStmt, _ = db.Prepare(SELECT)
	TGetStmt, _ = db.Prepare(TSELECT_ID)
}
