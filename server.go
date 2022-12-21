package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jubbyy/assessment/action"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	_ "github.com/lib/pq"
)

func main() {
	var rInit bool
	e := new(model.Expenses)
	e.Id = 5
	e.Title = "Title"
	e.Amount = 10
	e.Note = "Note"
	e.Tags = append(e.Tags, "a", "b")
	_ = e.Id
	rPort, PORTSet := os.LookupEnv("PORT")
	if !PORTSet {
		flPort := flag.String("port", "2565", "Running [PORT] for http server")
		rPort = *flPort
	}

	flInit, INITSet := os.LookupEnv("INIT2")
	if !INITSet {
		flag.BoolVar(&rInit, "init", true, "Initial Database")
		_ = rInit
	} else {
		var rInit, _ = strconv.ParseBool(flInit)
		_ = rInit
	}

	fmt.Println(os.Getenv("STATE"))
	fmt.Printf("rPort %v\n", rPort)
	fmt.Printf("rInit %v\n", rInit)

	action.DelExpense()
	action.PutExpense()

	connStr := os.Getenv("DATABASE_URL")
	database.ConnectDB(connStr)
	if rInit {
		fmt.Println("Initilizing Database ......")
		database.DB.Exec(database.DROP_TABLE)
		database.DB.Exec(database.CREATETABLE)
		database.MockData(10)
	}
	command := os.Getenv("ACTION")
	switch command {
	case "get":
		action.GetExpense(0)
	case "post":
		action.PostExpense(e)
	case "default":
		fmt.Println("Default Command")
	}

}
