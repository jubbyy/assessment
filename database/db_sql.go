package database

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/jubbyy/assessment/debug"
	"github.com/lib/pq"
)

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

type DBInterface struct {
	Db                                                      *sql.DB
	TGetstmt, Getstmt, Delstmt, Poststmt, Putstmt, Liststmt *sql.Stmt
}

var (
	TGetStmt, GetStmt, DelStmt, PostStmt, PutStmt, ListStmt *sql.Stmt
)

func ConnectDB(URL string) *sql.DB {

	debug.D("Openning DB Connection...")
	db, err := sql.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}
	db.Exec(CREATE_TABLE)

	debug.D("Database Connected and Table is ready.")
	GetStmt, _ = db.Prepare(SELECT_ID)
	DelStmt, _ = db.Prepare(DELETE_ID)
	PostStmt, _ = db.Prepare(INSERT)
	PutStmt, _ = db.Prepare(UPDATE_ID)
	ListStmt, _ = db.Prepare(SELECT)
	TGetStmt, _ = db.Prepare(TSELECT_ID)
	return db
}

func MockData(max int) {
	var cnum int
	var title, note string
	var tags []string

	debug.D("Mocking up data ....")
	//rand.Seed(30)
	for i := 1; i <= max; i++ {
		cnum = rand.Intn(1000)
		title = "Title:" + strconv.Itoa(cnum)
		note = "Notes:" + strconv.Itoa(cnum)
		tags = []string{"tag1", "tag2"}
		//		fmt.Println(cnum, title, note, tags)
		_, _ = PostStmt.Exec(title, cnum, note, pq.Array(&tags))
	}
	debug.D("Mocked up data ready")
}
