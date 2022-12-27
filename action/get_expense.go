package action

import (
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
)

func GetExpense(id int64) string {
	if id == 0 {
		debug.D("error no id to select")
		return "error / id is 0"
	}

	var e model.Expense
	var tags string
	err := database.GetStmt.QueryRow(id).Scan(&e.Id, &e.Title, &e.Amount, &e.Note, &tags)
	e.Tags = strings.Split(tags, ",")
	if err != nil {
		debug.D("Query Error")
		return "error / Query Error"
	}
	debug.D(e.AsJSON())
	return e.AsJSON()
}
