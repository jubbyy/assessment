package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	"github.com/lib/pq"
)

func PostExpense(c *gin.Context) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format JSON data (action.PostExpense)"})
		return
	}

	//	tags := strings.Join(e.Tags, ",")
	err := database.PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&e.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Insertion to database error"})
		return
	}
	c.JSON(http.StatusOK, e)
}
