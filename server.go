package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
	"github.com/jubbyy/assessment/myserver"
	_ "github.com/lib/pq"
)

var Config model.Configuration
var DB *sql.DB

func setup() {
	rel := flag.Bool("debugmode", false, "Run GIN as Debug mode - default false (releasemode)")
	init := flag.Bool("init", false, "Force Re-Initial Database")
	mock := flag.Bool("mock", false, "Create Mock Data (can't use without init) ")
	noweb := flag.Bool("noweb", false, "Do something backend without Web Service")
	log := flag.Bool("verboselog", false, "Enable Verbose/Debug Message")
	localhost := flag.Bool("localhost", false, "Running on Localhost Interface (for windows)")

	flag.Parse()

	Config.Init = *init
	Config.GinRelease = !*rel
	Config.Iface = ""
	Config.Port = "2565"
	Config.VerboseLog = *log
	Config.Mock = *mock
	Config.Noweb = *noweb

	oport := os.Getenv("PORT")
	_, oporterr := strconv.Atoi(oport)
	if oporterr == nil {
		Config.Port = oport
	}

	if *log {
		debug.Enabled = true
	}

	if *localhost {
		Config.Iface = "localhost"
	}

	debug.D(fmt.Sprintf("%v", Config))
}

func main() {

	setup()

	URL := os.Getenv("DATABASE_URL")
	DB = database.ConnectDB(URL)

	if !Config.Noweb {
		router := myserver.StartAndRoute(Config.GinRelease)
		router.Run(Config.Iface + ":" + Config.Port)
	}

	if Config.Init {
		debug.D("Initialising....")
		if DB == nil {
			panic("lost databaes ?")
		}
		DB.Exec(database.DROP_TABLE)
		DB.Exec(database.CREATE_TABLE)
		if Config.Mock {
			database.MockData(10)
		}
	}
}
