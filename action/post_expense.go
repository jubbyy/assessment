package action

import (
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func PostExpense(e model.Expense) {
	st, err := database.DB.Prepare(database.INSERT)
	HasErr(500, err)
	defer st.Close() // Prepared statements take up server resources and should be closed after use.

	tags := strings.Join(e.Tags, ",")
	res, err := st.Exec(e.Id, e.Title, e.Amount, e.Note, tags)
	HasErr(500, err)

	newid, err := res.LastInsertId()
	HasErr(400, err)
	e.Id = newid

}
