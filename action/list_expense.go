package action

import (
	"fmt"
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
)

func ListExpense() {
	st, err := database.DB.Prepare(database.SELECT_LIMIT)
	if err != nil {
		panic(err)
	}

	rows, err := st.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var results []model.Expenses
	var id int64
	var title, note, tags string
	var amount float32
	var result model.Expenses

	for rows.Next() {
		err := rows.Scan(&id, &title, &amount, &note, &tags)
		if err != nil {
			panic(err)
		}

		result.Id = id
		result.Title = title
		result.Amount = amount
		result.Note = note
		result.Tags = strings.Split(tags, ",")
		results = append(results, result)

	}
	debug.D(fmt.Sprintf("%v", results))
}
