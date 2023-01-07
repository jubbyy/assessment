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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		id = 0
	}

	res, err := database.PutStmt.Exec(id, e.Title, e.Amount, e.Note, pq.Array(&e.Tags))

	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id : " + c.Param("id") + " not found."})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "invalid data (action.PutExpense)"})
		return
	}

	e.Id = id
	c.JSON(http.StatusOK, e)

}

func UpdateExpense(c *gin.Context, sg *database.StatementGroup) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data (action.UpdateExpense)"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		id = 0
	}

	res, er := sg.PutStmt.Exec(id, e.Title, e.Amount, e.Note, pq.Array(&e.Tags))
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update error in action.UpdateExpense"})
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id : " + c.Param("id") + " not found."})
		return
	}

	e.Id = id
	c.JSON(http.StatusOK, e)

}
