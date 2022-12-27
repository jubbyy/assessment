package action

import (
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func PostExpense(e model.Expense) int64 {
	st, err := database.DB.Prepare(database.INSERT)
	HasErr(504, err)
	defer st.Close() // Prepared statements take up server resources and should be closed after use.

	tags := strings.Join(e.Tags, ",")
	var newid int64
	err = st.QueryRow(e.Title, e.Amount, e.Note, tags).Scan(&newid)
	HasErr(500, err)
	return newid
}
