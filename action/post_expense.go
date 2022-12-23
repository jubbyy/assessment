package action

import (
	"fmt"
	"log"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func PostExpense(e model.Expense) {
	fmt.Printf("%v", e)
	st, err := database.DB.Prepare(database.INSERT)
	if err != nil {
		fmt.Printf("%T", database.DB)
		log.Fatal(err)
	}
	defer st.Close() // Prepared statements take up server resources and should be closed after use.
	st.Exec(e.Title, e.Amount, e.Note, e.Tags)
}
