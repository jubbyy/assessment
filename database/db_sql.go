package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
)

var DB *sql.DB
var CREATETABLE = `CREATE TABLE  IF NOT EXISTS expenses (
	id serial PRIMARY KEY,
	title VARCHAR ( 140 ) NOT NULL,
	amount int  NOT NULL,
	note VARCHAR ( 255 ),
    tags VARCHAR (255)
)`
var DROP_TABLE = `drop table expenses`

var SELECT = `select * from expenses`
var SELECT_ID = `select * from expenses where id = ? `
var DELETE_ID = `delete from expenses where id = ?`
var UPDATE_ID = `update `
var INSERT = `insert into expenses (title,amount,note,tags) values(?,?,?,?)`
var MOCK_RECORD = `insert into expenses (title,amount,note,tags) values('Test Expenses',501,'Mock Record','tags1,tags2,tags3')`

func ConnectDB(URL string) {
	db, err := sql.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}

	DB = db
}
func MockData(max int) {
	var cnum int
	var title, note, tags string
	mockst := `insert into expenses(title,amount,note,tags) values($1, $2 ,$3,$4)`

	st, err := DB.Prepare(mockst)
	if err != nil {
		fmt.Println(err)
		panic("Prepare Error")
	}

	rand.Seed(30)
	for i := 1; i < max; i++ {
		cnum = rand.Intn(1000)
		fmt.Println(cnum)
		title = "Title:" + strconv.Itoa(cnum)
		note = "Notes:" + strconv.Itoa(cnum)
		tags = "tag1,tag" + strconv.Itoa(cnum)
		//		fmt.Println(cnum, title, note, tags)
		_, _ = st.Exec(title, cnum, note, tags)
	}
}
