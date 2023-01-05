//go:build unittest

package action

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	"github.com/jubbyy/assessment/database"
	"net/http/httptest"
	"testing"
)

var DB *sql.DB

func init() {
	db, mock, err := sqlmock.New()
	_ = mock
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db.Begin()
	database.GetStmt, _ = db.Prepare(database.SELECT_ID)
	database.DelStmt, _ = db.Prepare(database.DELETE_ID)
	database.PostStmt, _ = db.Prepare(database.INSERT)
	database.PutStmt, _ = db.Prepare(database.UPDATE_ID)
	database.ListStmt, _ = db.Prepare(database.SELECT)
	database.TGetStmt, _ = db.Prepare(database.TSELECT_ID)
	DB = db
}
func TestGetExpense(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}

	GetExpense(c)
	if w.Code != 200 {
		t.Error(w.Code, "")
	}
}
