package myserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/action"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/myusers"
)

var (
	Ginusers = gin.Accounts{myusers.Good.Name: myusers.Good.Password,
		myusers.Admin.Name: myusers.Admin.Password}
)

func StartAndRoute(releasemode bool, DH *database.DatabaseHandler) *gin.Engine {
	if releasemode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "KBTG Pong"})
	})
	router.GET("/exp", action.ListExpense)
	router.GET("/exp/:id", action.GetExpense)

	authen := router.Group("/", gin.BasicAuth(Ginusers))
	authen.GET("/expenses", action.ListExpense)
	authen.GET("/expenses/:id", action.GetExpense)
	authen.POST("/expenses", func(c *gin.Context) {
		action.PostExpense(c, DH)
	})
	authen.PUT("/expenses/:id", action.PutExpense)
	authen.DELETE("/expenses/:id", action.DelExpense)

	return router
}

func RouteSetup(releasemode bool, sg *database.StatementGroup) *gin.Engine {
	if releasemode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "KBTG Pong"})
	})
	router.GET("/exp", action.ListExpense)
	router.GET("/exp/:id", action.GetExpense)

	authen := router.Group("/", gin.BasicAuth(Ginusers))
	authen.GET("/expenses", func(c *gin.Context) {
		action.ListAllExpense(c, sg)
	})
	authen.GET("/expenses/:id", func(c *gin.Context) {
		action.GetAnExpense(c, sg)
	})
	authen.POST("/expenses", func(c *gin.Context) {
		action.AddExpense(c, sg)
	})
	authen.PUT("/expenses/:id", func(c *gin.Context) {
		action.UpdateExpense(c, sg)
	})
	authen.DELETE("/expenses/:id", func(c *gin.Context) {
		action.DeleteExpense(c, sg)
	})
	authen.DELETE("/expenses/reset", func(c *gin.Context) {
		action.ResetExpense(c, sg)
	})
	return router
}
