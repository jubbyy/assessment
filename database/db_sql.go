package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/jubbyy/assessment/debug"
	"github.com/jubbyy/assessment/model"
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

type DatabaseHandler struct {
	Db                                                                              *sql.DB
	DropStmt, CreateStmt, OneGetStmt, GetStmt, DelStmt, PostStmt, PutStmt, ListStmt *sql.Stmt
	Message                                                                         string
	Errcode                                                                         int
}
type StatementGroup struct {
	DropStmt, CreateStmt, OneGetStmt, GetStmt, DelStmt, PostStmt, PutStmt, ListStmt *sql.Stmt
}

var (
	TGetStmt, DelStmt, PostStmt, PutStmt, ListStmt *sql.Stmt
)

func (dh *DatabaseHandler) ConnectDB(URL string) {
	var err error
	debug.D("Openning DB Connection...")
	//	dh.Db, _ = sql.Open("postgres", URL)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(URL)

	debug.D("Database Connected and Table is ready.")
	dh.CreateStmt, _ = db.Prepare(CREATE_TABLE)
	dh.DropStmt, _ = db.Prepare(DROP_TABLE)
	dh.GetStmt, _ = db.Prepare(SELECT_ID)
	dh.DelStmt, _ = dh.Db.Prepare(DELETE_ID)
	dh.PostStmt, _ = dh.Db.Prepare(INSERT)
	dh.PutStmt, _ = dh.Db.Prepare(UPDATE_ID)
	dh.ListStmt, _ = dh.Db.Prepare(SELECT)
	dh.OneGetStmt, _ = dh.Db.Prepare(TSELECT_ID)
	dh.Message = "Ready"
	dh.Errcode = 208

	_, er := dh.CreateStmt.Exec()
	if er != nil {
		panic(er.Error())
	}
}
func DBConnect(URL string) (db *sql.DB) {
	db, err := sql.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}
	return
}

func DBControl(URL string) (*sql.DB, *StatementGroup) {
	db, err := sql.Open("postgres", URL)

	if err != nil {
		panic(err.Error())
	}
	var stmt *sql.Stmt
	p := func(q string) *sql.Stmt {
		stmt, _ = db.Prepare(q)
		return stmt
	}
	if err != nil {
		panic(err.Error())
	}

	res, er := db.Exec(CREATE_TABLE)
	_ = res
	if er != nil {
		log.Println(er.Error())
	}
	return db, &StatementGroup{
		DropStmt:   p(DROP_TABLE),
		CreateStmt: p(CREATE_TABLE),
		OneGetStmt: p(TSELECT_ID),
		GetStmt:    p(SELECT_ID),
		DelStmt:    p(DELETE_ID),
		PostStmt:   p(INSERT),
		PutStmt:    p(UPDATE_ID),
		ListStmt:   p(SELECT),
	}
}

func (dh *DatabaseHandler) Create(e model.Expense) (code int, eout model.Expense) {
	eout = e
	//	return dh.Errcode, eout
	err := dh.PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&eout.Id)

	if err != nil {
		return http.StatusBadRequest, e
	}

	return http.StatusOK, eout
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
