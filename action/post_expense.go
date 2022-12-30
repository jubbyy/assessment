package action

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func PostExpense(c *gin.Context) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format JSON data (action.PostExpense)"})
		return
	}

	tags := strings.Join(e.Tags, ",")
	var newid int64
	err := database.PostStmt.QueryRow(e.Title, e.Amount, e.Note, tags).Scan(&newid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Insertion to database error"})
		return
	}
	e.Id = newid
	c.JSON(http.StatusOK, e)
}
