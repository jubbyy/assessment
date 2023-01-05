package action

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	"github.com/lib/pq"
)

func GetExpense(c *gin.Context) {
	var e model.Expense
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		id = 0
	}
	err := database.TGetStmt.QueryRow(id).Scan(&e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "expense id : " + c.Param("id") + " not Found"})
		return
	}
	e.Id = id

	c.JSON(http.StatusOK, e)
}
