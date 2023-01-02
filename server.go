package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
	"github.com/jubbyy/assessment/myserver"
	_ "github.com/lib/pq"
)

var Config model.Configuration

func setup() {
	rel := flag.Bool("debugmode", false, "Run GIN as Debug mode - default false (releasemode)")
	init := flag.Bool("init", false, "Force Re-Initial Database")
	mock := flag.Bool("mock", false, "Create Mock Data (can't use without init) ")
	port := flag.String("port", "2565", "Service Port")
	action := flag.String("action", "get", "Action=web start web server")
	log := flag.Bool("verboselog", false, "Enable Verbose/Debug Message")

	flag.Parse()

	Config.Init = *init
	Config.GinRelease = !*rel
	Config.Iface = "localhost"
	Config.Port = *port
	Config.Action = strings.ToLower(*action)
	Config.VerboseLog = *log
	Config.Mock = *mock

	oport := os.Getenv("PORT")
	_, oporterr := strconv.Atoi(oport)
	if oporterr == nil {
		Config.Port = oport
	}

	if *log {
		debug.Enabled = true
	}

	debug.D(fmt.Sprintf("%v", Config))
}

func main() {

	setup()

	URL := os.Getenv("DATABASE_URL")
	database.ConnectDB(URL)

	if Config.Action == "web" {
		router := myserver.StartAndRoute(Config.GinRelease)
		router.Run(Config.Iface + ":" + Config.Port)
	}

	//	var e model.Expense
	//	database.TGetStmt.QueryRow(1).Scan(&e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
	//	fmt.Println(e)

	if Config.Init {
		debug.D("Initialising....")
		database.DB.Exec(database.DROP_TABLE)
		database.DB.Exec(database.CREATE_TABLE)
		if Config.Mock {
			database.MockData(10)
		}
	}
}
