package action

import (
	"log"

	"github.com/jubbyy/assessment/database"
)

func GetExpense(id int64) string {
	if id == 0 {
		log.Fatal("error no id to select")
		return "error"
	}
	st, err := database.DB.Prepare(database.SELECT_ID)
	if err != nil {
		log.Fatal(err)
	}

	//    err := stmt.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.Quantity)
	defer st.Close()
	st.Query(id)
	return "success"
}
