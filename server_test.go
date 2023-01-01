package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/myserver"
	"github.com/jubbyy/assessment/myusers"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func TestConnectDB(t *testing.T) {
	//func SetupDatabase(t *testing.T) {
	URL := os.Getenv("TESTDATABASE_URL")
	if URL == "" {
		t.Log("WARNING!!.. Testing Database Connection.... DATBASE_URL ")
		os.Getenv("DATABASE_URL")
	} else {
		t.Log("Connecting to Database .... TESTDATBASE_URL ")
	}
	database.ConnectDB(URL)
	t.Log("Droping Table and Create a new one")
	database.DB.Exec(database.DROP_TABLE)
	database.DB.Exec(database.CREATETABLE)
	//database.DB.Exec(database.TEST_RECORD1)
	//database.DB.Exec(database.TEST_RECORD2)
	//database.DB.Exec(database.TEST_RECORD3)
}

func TestStartGinWebServer(t *testing.T) {
	t.Log("Starting Gin Web Server")
	GinRelease := true
	r = myserver.StartAndRoute(GinRelease)
}

func TestPing(t *testing.T) {
	expectResponse := `{"message":"KBTG Pong"}`

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoryExp01(t *testing.T) {
	jsonRequest := `{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`
	expectResponse := `{"id":1,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer([]byte(jsonRequest)))
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoryExp02(t *testing.T) {
	expectResponse := `{"id":1,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/expenses/1", nil)
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestStoryExp03(t *testing.T) {
	jsonRequest := `{
		"title": "apple smoothie",
		"amount": 89,
		"note": "no discount", 
		"tags": ["beverage"]
	}`
	expectResponse := `{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}`
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/expenses/1", bytes.NewBuffer([]byte(jsonRequest)))
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestStoryExp04(t *testing.T) {
	jsonR1 := `{
		"title": "apple smoothie",
		"amount": 89,
		"note": "no discount",
		"tags": ["beverage"]
	}`
	jsonR2 := `{
		"title": "iPhone 14 Pro Max 1TB",
		"amount": 66900,
		"note": "birthday gift from my love", 
		"tags": ["gadget"]
	}`
	expectResponse := `[{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]},{"id":2,"title":"iPhone 14 Pro Max 1TB","amount":66900,"note":"birthday gift from my love","tags":["gadget"]}]`

	t.Log("Dropping old Data before starting testing Exp04")
	database.DB.Exec(database.DROP_TABLE)
	database.DB.Exec(database.CREATETABLE)

	w := httptest.NewRecorder()
	wt := httptest.NewRecorder()

	t.Log("Mocking 2 records")
	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer([]byte(jsonR1)))
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(wt, req)
	req, _ = http.NewRequest("POST", "/expenses", bytes.NewBuffer([]byte(jsonR2)))
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(wt, req)

	req, _ = http.NewRequest("GET", "/expenses", nil)
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBadAuthentication(t *testing.T) {
	t.Log("Testing bad username")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/expenses", nil)
	req.SetBasicAuth(myusers.Bad.Name, myusers.Bad.Password)
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusUnauthorized {
		t.Errorf("Must ask for password - return statuscode %v", status)
	}
}
