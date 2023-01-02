package action

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	"github.com/lib/pq"
)

func PutExpense(c *gin.Context) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input data (action.PutExpense)"})
		return
	}

	e.Id, _ = strconv.Atoi(c.Param("id"))
	res, err := database.PutStmt.Exec(e.Id, e.Title, e.Amount, e.Note, pq.Array(&e.Tags))
	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": c.Param("id") + " (id) not found."})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid data (action.PutExpense)"})
		return
	}
	c.JSON(http.StatusOK, e)
}
