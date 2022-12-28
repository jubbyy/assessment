package action

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
)

func GetExpense(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var e model.Expense
	var tags string
	err := database.GetStmt.QueryRow(id).Scan(&e.Id, &e.Title, &e.Amount, &e.Note, &tags)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID ( " + strconv.FormatInt(id, 10) + " ) not found. "})
		return
	}
	e.Tags = strings.Split(tags, ",")
	c.IndentedJSON(http.StatusOK, e)
}
