package action

import (
	"log"
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
)

func PutExpense(e model.Expense) bool {
	if e.Id == 0 {
		debug.D("ID 0 is not allow - " + e.AsJSON())
		return false
	}
	log.Println(e.AsJSON())
	tags := strings.Join(e.Tags, ",")
	_, err := database.PutStmt.Exec(e.Id, e.Title, e.Amount, e.Note, tags)
	HasErr(500, err)
	return true

}
