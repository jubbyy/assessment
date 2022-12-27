package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jubbyy/assessment/debug"
)

var DB *sql.DB
var CREATETABLE = `CREATE TABLE  IF NOT EXISTS expenses (
	id serial PRIMARY KEY,
	title VARCHAR ( 140 ) NOT NULL,
	amount float  NOT NULL,
	note VARCHAR ( 255 ),
    tags VARCHAR (255)
)`
var DROP_TABLE = `drop table expenses`
var SELECT_LIMIT = `select * from expenses limit 10`
var SELECT_ID = `select id,title,amount,note,tags from expenses where id = $1 `
var DELETE_ID = `delete from expenses where id = $1`
var UPDATE_ID = `update expenses set title=$2, amount=$3, note = $4, tags = $5 where id=$1`
var INSERT = `insert into expenses (title,amount,note,tags) values($1,$2,$3,$4) RETURNING id`
var MOCK_RECORD = `insert into expenses (title,amount,note,tags) values('Test Expenses',501,'Mock Record','tags1,tags2,tags3')`

var (
	GetStmt, DelStmt, PostStmt, PutStmt *sql.Stmt
)

func ConnectDB(URL string) {
	debug.D("Openning DB Connection...")
	db, err := sql.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}
	db.Exec(CREATETABLE)
	debug.D("Database Connected and Table is ready.")
	GetStmt, err = db.Prepare(SELECT_ID)
	DelStmt, err = db.Prepare(DELETE_ID)
	PostStmt, err = db.Prepare(INSERT)
	PutStmt, err = db.Prepare(UPDATE_ID)

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
	debug.D("Mocking up data ....")
	//rand.Seed(30)
	for i := 1; i <= max; i++ {
		cnum = rand.Intn(1000)
		title = "Title:" + strconv.Itoa(cnum)
		note = "Notes:" + strconv.Itoa(cnum)
		tags = "tag1,tag" + strconv.Itoa(cnum)
		//		fmt.Println(cnum, title, note, tags)
		_, _ = st.Exec(title, cnum, note, tags)
	}
	debug.D("Mocked up data ready")
}
