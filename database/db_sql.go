package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jubbyy/assessment/debug"
)

var mockResponse = `{"id":1,"title":"Title1","amount":1111.11,"notes":"notes1","tags":["tags1","tags2"]}`
var DB *sql.DB

var (
	CREATETABLE = `CREATE TABLE  IF NOT EXISTS expenses (
		id serial PRIMARY KEY,
		title VARCHAR ( 140 ) NOT NULL,
		amount float  NOT NULL,
		note VARCHAR ( 255 ),
		tags VARCHAR (255)
	)`

	CREATE_TABLE = `	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`
	DROP_TABLE   = `drop table expenses`
	SELECT       = `select * from expenses order by id`
	SELECT_ID    = `select id,title,amount,note,tags from expenses where id = $1`
	DELETE_ID    = `delete from expenses where id = $1`
	UPDATE_ID    = `update expenses set title=$2, amount=$3, note = $4, tags = $5 where id=$1`
	INSERT       = `insert into expenses (title,amount,note,tags) values($1,$2,$3,$4) RETURNING id`
	TEST_RECORD1 = `insert into expenses (title,amount,note,tags) values('Title1',1111.11,'Note1','tags1,tags2')`
	TEST_RECORD2 = `insert into expenses (title,amount,note,tags) values('Title2',2222.22,'Note2','tags2,tags3')`
	TEST_RECORD3 = `insert into expenses (title,amount,note,tags) values('Title3',3333.33,'Note3','tags1,tags3')`
)

var (
	GetStmt, DelStmt, PostStmt, PutStmt, ListStmt *sql.Stmt
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
	ListStmt, err = db.Prepare(SELECT)

	DB = db
}

func MockDataNew() {
	DB.Exec(DROP_TABLE)
	DB.Exec(CREATETABLE)
	DB.Exec(TEST_RECORD1)
	DB.Exec(TEST_RECORD2)
	DB.Exec(TEST_RECORD3)
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
