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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something wrong in action.ListExpense"})
		return
	}
	defer rows.Close()

	var results []model.Expense
	var result model.Expense

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid rows conversion action.ListExpense"})
			return
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		e := [0]string{}
		c.JSON(http.StatusOK, e)
		return
	}

	c.JSON(http.StatusOK, results)
}
func ListAllExpense(c *gin.Context, sg *database.StatementGroup) {

	rows, err := sg.ListStmt.Query()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something wrong in action.ListAllExpense"})
		return
	}
	defer rows.Close()

	var results []model.Expense
	var result model.Expense

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid rows conversion action.ListAllExpense"})
			return
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		e := [0]string{}
		c.JSON(http.StatusOK, e)
		return
	}

	c.JSON(http.StatusOK, results)
}
