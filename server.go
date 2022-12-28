package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/action"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
	_ "github.com/lib/pq"
)

var Config model.Configuration

func setup() {
	deb := flag.Bool("debug", false, "Enable debugging message")
	init := flag.Bool("init", false, "Force Re-Initial Database")
	port := flag.String("port", "2565", "Service Port")
	action := flag.String("action", "get", "Action for test can be get put post delete")

	flag.Parse()

	Config.Init = *init
	Config.Debug = *deb
	Config.Iface = "localhost"
	Config.Port = *port
	Config.Action = strings.ToLower(*action)

	oport := os.Getenv("PORT")
	_, oporterr := strconv.Atoi(oport)
	if oporterr == nil {
		Config.Port = oport
	}
	debug.D(fmt.Sprintf("%v", Config))
}

func webserver() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	router.GET("/expenses", action.ListExpense)

	router.GET("/expenses/:id", action.GetExpense)

	router.POST("/expenses", func(c *gin.Context) {
		var e model.Expense
		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := action.PostExpense(e)
		e.Id = res
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, e.AsJSON())
	})

	router.DELETE("/expenses/:id", func(c *gin.Context) {
		id := c.Param("id")
		nid, _ := strconv.ParseInt(id, 10, 64)
		res, _ := database.DelStmt.Exec(nid)
		if numrow, _ := res.RowsAffected(); numrow != 1 {
			c.JSON(http.StatusNotFound, gin.H{"message": id + " not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": id + " was deleted."})
	})

	router.PUT("/expenses/:id", func(c *gin.Context) {
		var json model.Expense
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		json.Id = id
		if action.PutExpense(json) {
			c.JSON(http.StatusOK, gin.H{"message": "record id : " + c.Param("id") + " was updated"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": c.Param("id") + " not FOUND"})
		}
	})

	router.Run(Config.Iface + ":" + Config.Port)
}

func main() {
	e := model.Expense{50, "title E", 50.0, "Note E", []string{"Tags A", "Tags E"}}

	setup()
	//webserver()
	//fmt.Printf("\n%v", rInit)
	//	fmt.Printf("rPort %v\n", rPort)
	//	fmt.Printf("rInit %v\n", rInit)
	//	action.PutExpense()

	connStr := os.Getenv("DATABASE_URL")
	database.ConnectDB(connStr)
	if Config.Init {
		debug.D("Initialising....")
		database.DB.Exec(database.DROP_TABLE)
		database.DB.Exec(database.CREATETABLE)
		database.MockData(10)
	}

	switch Config.Action {
	case "post":
		action.PostExpense(e)
	case "web":
		webserver()
	default:
		fmt.Println("Default Command")
	}

}
