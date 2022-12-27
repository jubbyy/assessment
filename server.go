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
	//	gin.SetMode(gin.ReleaseMode)

	//	router.GET("/someGet", getting)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	router.GET("/expense/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		if id == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "id 0 not found."})
			panic("wrong id")
		}
		e := action.GetExpense(id)
		_ = e
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, e)
	})

	router.POST("/expense", func(c *gin.Context) {
		var json model.Expense
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := action.PostExpense(json)
		json.Id = res
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, json.AsJSON())
	})

	router.DELETE("/expense/:id", func(c *gin.Context) {
		id := c.Param("id")
		inid, _ := strconv.Atoi(id)

		if inid == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": id + " not found "})
			return
		}

		st, err := database.DB.Prepare(database.DELETE_ID)
		action.HasErr(500, err)

		res, err := st.Exec(id)
		if numrow, _ := res.RowsAffected(); numrow != 1 {
			c.JSON(http.StatusNotFound, gin.H{"message": id + "not found"})
		}

		c.JSON(http.StatusOK, gin.H{"message": id + " was deleted."})
	})

	router.PUT("/expense/:id", func(c *gin.Context) {
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
	case "get":
		action.GetExpense(1)
	case "post":
		action.PostExpense(e)
	case "delete":
		action.DelExpense(1)
		//	case "update":
		//		action.PutExpense(`{"title":"Test","notes":"Notes","tags":["tagsss","tags2"],"amount":55.0}`)
	case "list":
		action.ListExpense()
	case "web":
		webserver()
	default:
		fmt.Println("Default Command")
	}

}
