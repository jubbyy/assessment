package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	"github.com/lib/pq"
)

func ListExpense(c *gin.Context) {

	rows, err := database.ListStmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var results []model.Expense
	var result model.Expense

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid rows data"})
			return
		}
		results = append(results, result)
	}
	c.JSON(http.StatusOK, results)
}
