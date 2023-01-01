package myserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/action"
	"github.com/jubbyy/assessment/myusers"
)

var (
	Ginusers = gin.Accounts{myusers.Good.Name: myusers.Good.Password,
		myusers.Admin.Name: myusers.Admin.Password}
)

func StartAndRoute(releasemode bool) *gin.Engine {
	if releasemode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "KBTG Pong"})
	})

	authen := router.Group("/", gin.BasicAuth(Ginusers))
	authen.GET("/expenses", action.ListExpense)
	authen.GET("/expenses/:id", action.GetExpense)
	authen.POST("/expenses", action.PostExpense)
	authen.PUT("/expenses/:id", action.PutExpense)
	authen.DELETE("/expenses/:id", action.DelExpense)

	return router
}
