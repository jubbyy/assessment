package action

import (
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func PutExpense(jsonstring string) {
	e := model.NewExpenseFromJSON(jsonstring)
	//	var e model.Expense
	if e.Id == 0 {
		HasErr(400, "error")
	}

	st, err := database.DB.Prepare(database.UPDATE_ID)
	HasErr(500, err)
	defer st.Close()

	tags := strings.Join(e.Tags, ",")
	res, err := st.Exec(e.Id, e.Title, e.Amount, e.Note, tags)
	HasErr(500, err)

	newid, err := res.LastInsertId()
	HasErr(500, err)
	e.Id = newid

}
