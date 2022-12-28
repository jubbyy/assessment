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

func construct() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	router.GET("/expenses", action.ListExpense)

	router.GET("/expenses/:id", action.GetExpense)

	router.POST("/expenses", action.PostExpense)

	router.DELETE("/expenses/:id", action.DelExpense)

	router.PUT("/expenses/:id", action.PutExpense)

	router.Run(Config.Iface + ":" + Config.Port)
}

func main() {

	setup()
	//construct()
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

	if Config.Action == "web" {
		construct()
	}

}
