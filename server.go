package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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
	noweb := flag.Bool("noweb", false, "Do something backend without Web Service")
	log := flag.Bool("verboselog", false, "Enable Verbose/Debug Message")
	localhost := flag.Bool("localhost", false, "Running only on Localhost Interface ")

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

	var srv *http.Server
	log.Println("Spinning up...... ")
	setup()

	URL := os.Getenv("DATABASE_URL")
	log.Println(URL)
	db, sg := database.DBControl(URL)
	sg.CreateStmt.Exec()

	defer CloseDB(db)

	if !Config.Noweb {

		router := myserver.RouteSetup(Config.GinRelease, sg)
		srv = &http.Server{
			Addr:    Config.Iface + ":" + Config.Port,
			Handler: router,
		}

		go func() {
			log.Println("Listening on " + Config.Iface + ":" + Config.Port)

			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

	}
	if Config.Init {
		debug.D("Initialising....")
		if db == nil {
			panic("lost databaes ?")
		}
		sg.DropStmt.Exec()
		sg.CreateStmt.Exec()

		if Config.Mock {
			database.MockData(10)
		}
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Shutting down.....")
}

func CloseDB(db io.Closer) {
	db.Close()
}
