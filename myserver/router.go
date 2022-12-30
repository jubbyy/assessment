package myserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/action"
)

func StartAndRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello KTBG Go01"})
	})

	router.GET("/expenses", action.ListExpense)
	router.GET("/expenses/:id", action.GetExpense)

	router.POST("/expenses", action.PostExpense)

	router.DELETE("/expenses/:id", action.DelExpense)

	router.PUT("/expenses/:id", action.PutExpense)

	return router
}
