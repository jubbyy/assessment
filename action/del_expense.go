package action

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
)

func DelExpense(c *gin.Context) {
	id := c.Param("id")
	nid, _ := strconv.ParseInt(id, 10, 64)
	res, _ := database.DelStmt.Exec(nid)
	if numrow, _ := res.RowsAffected(); numrow != 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": id + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": id + " was deleted."})

}
