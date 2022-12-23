package action

import (
	"fmt"

	"github.com/jubbyy/assessment/database"
)

func DelExpense(id int) {
	st, err := database.DB.Prepare(database.DELETE_ID)

	if err != nil {
		panic(err)
	}

	res, err := st.Exec(id)
	if err != nil {
		panic(err)
	}

	afrow, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID %v deleted successfully with %v affected\n", id, afrow)

}
