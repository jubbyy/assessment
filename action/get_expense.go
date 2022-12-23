package action

import (
	"fmt"
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
)

func GetExpense(id int64) string {
	if id == 0 {
		debug.D("error no id to select")
		return "error"
	}
	st, err := database.DB.Prepare(database.SELECT_ID)
	if err != nil {
		debug.D(err.Error())
	}
	var e model.Expense
	var tags string
	er := st.QueryRow(id).Scan(&e.Id, &e.Title, &e.Amount, &e.Note, &tags)
	e.Tags = strings.Split(tags, ",")
	if er != nil {
		debug.D("Query Error")
	}
	defer st.Close()
	fmt.Println(e.AsJSON())
	return "success"
}
