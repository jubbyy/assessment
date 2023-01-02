package action

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	"github.com/lib/pq"
)

/*
	func GetExpense(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		//	var e model.Expense
		//var tags []sql.NullString
		var tags string

		var eid int
		var title, note string
		var amount float32
		//	var tags string
		err := database.GetStmt.QueryRow(id).Scan(&eid, &title, &amount, &note, &tags)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, tags)
	}
*/
func GetExpense(c *gin.Context) {
	var e model.Expense
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		id = 0
	}

	err := database.TGetStmt.QueryRow(id).Scan(&e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, e)
}
