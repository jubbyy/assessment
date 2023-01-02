package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jubbyy/assessment/debug"
	"github.com/lib/pq"
)

var mockResponse = `{"id":1,"title":"Title1","amount":1111.11,"notes":"notes1","tags":["tags1","tags2"]}`
var DB *sql.DB

var (
	CREATE_TABLE = `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`
	DROP_TABLE = `drop table expenses`
	SELECT     = `select id,title,amount,note,tags from expenses order by id`
	SELECT_ID  = `select id,title,amount,note,tags from expenses where id = $1`
	TSELECT_ID = `select title,amount,note,tags from expenses where id = $1`
	DELETE_ID  = `delete from expenses where id = $1`
	UPDATE_ID  = `update expenses set title=$2, amount=$3, note = $4, tags = $5 where id=$1`
	INSERT     = `insert into expenses (title,amount,note,tags) values($1,$2,$3,$4) RETURNING id`
)

var (
	TGetStmt, GetStmt, DelStmt, PostStmt, PutStmt, ListStmt *sql.Stmt
)

func ConnectDB(URL string) {

	debug.D("Openning DB Connection...")
	DB, err := sql.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}
	DB.Exec(CREATE_TABLE)

	debug.D("Database Connected and Table is ready.")
	GetStmt, _ = DB.Prepare(SELECT_ID)
	DelStmt, _ = DB.Prepare(DELETE_ID)
	PostStmt, _ = DB.Prepare(INSERT)
	PutStmt, _ = DB.Prepare(UPDATE_ID)
	ListStmt, _ = DB.Prepare(SELECT)
	fmt.Println(DB.Stats())

}

func MockData(max int) {
	var cnum int
	var title, note string
	var tags []string
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
		tags = []string{"tag1", "tag2"}
		//		fmt.Println(cnum, title, note, tags)
		_, _ = st.Exec(title, cnum, note, pq.Array(&tags))
	}
	debug.D("Mocked up data ready")
}
