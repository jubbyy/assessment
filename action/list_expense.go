package action

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func ListExpense(c *gin.Context) {

	rows, err := database.ListStmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var results []model.Expense
	var id int64
	var title, note, tags string
	var amount float32
	var result model.Expense

	for rows.Next() {
		err := rows.Scan(&id, &title, &amount, &note, &tags)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid rows data"})
			return
		}

		result.Id = id
		result.Title = title
		result.Amount = amount
		result.Note = note
		result.Tags = strings.Split(tags, ",")
		results = append(results, result)
	}
	c.JSON(http.StatusOK, results)
}
