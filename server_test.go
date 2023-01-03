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
	t.Log("Connecting to : " + URL)
	DB = database.ConnectDB(URL)

	t.Log("Droping Table and Create a new one")
	res, err := DB.Exec(database.DROP_TABLE)
	_ = res
	if err != nil {
		t.Log("Can't Drop Table")
		t.Log(err.Error())
	}
	res, err = DB.Exec(database.CREATE_TABLE)
	_ = res
	if err != nil {
		t.Log("Can't Create Table")
		t.Log(err.Error())
	}

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
func TestStoryExp02NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/expenses/1000", nil)
	req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestStoryExp03(t *testing.T) {
	jsonRequest := `{
		"title": "apple smoothie",
		"amount": 69,
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
	DB.Exec(database.DROP_TABLE)
	DB.Exec(database.CREATE_TABLE)

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
