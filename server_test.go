//go:build integration

package main

import (
	"bytes"
	//	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/myserver"
	"github.com/jubbyy/assessment/myusers"
	pq "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

/*func init() {
	URL := os.Getenv("DATABASE_URL")
	if os.Getenv("TESTDATABASE_URL") != "" {
		fmt.Println("Connecting to Database .... TESTDATBASE_URL ")
		URL = os.Getenv("TESTDATABSE_URL")
	} else {
		fmt.Println("!!WARNING!!.. Testing Database Connection..to Prodcution DATBASE_URL ")
	}
	fmt.Println("Connecting to %s", URL)
	fmt.Println("!!WARNING!! Deleting all DATA")
	DB = database.ConnectDB(URL)
	DB.Exec(database.DROP_TABLE)
	DB.Exec(database.CREATE_TABLE)
}*/

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
	t.Log("Test Case : Web Server /ping Ready?")
	expectResponse := `{"message":"KBTG Pong"}`

	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoryExp01(t *testing.T) {
	t.Run("Create an expense - should return id = 1 ", func(t *testing.T) {
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

	})
}

func TestStoryExp02(t *testing.T) {
	t.Run("Get expense id=1 -SUCCESS-", func(t *testing.T) {

		expectResponse := `{"id":1,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}`
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/expenses/1", nil)
		req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
		r.ServeHTTP(w, req)

		responseData, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, expectResponse, string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("Get expense id 0 -404 NOT FOUND-", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/expenses/0", nil)
		req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

	})
}
func TestStoryExp03(t *testing.T) {
	jsonRequest := `{
		"title": "apple smoothie",
		"amount": 89,
		"note": "no discount", 
		"tags": ["beverage"]
	}`
	t.Run("Update expense id=1 -SUCCESS-", func(t *testing.T) {

		expectResponse := `{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}`
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("PUT", "/expenses/1", bytes.NewBuffer([]byte(jsonRequest)))
		req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
		r.ServeHTTP(w, req)

		responseData, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, expectResponse, string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("Update to id=0 -404 NOT FOUND-", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("PUT", "/expenses/0", bytes.NewBuffer([]byte(jsonRequest)))
		req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

	})

}

func TestStoryExp04(t *testing.T) {
	t.Log("Dropping old DATA for testing Exp04")
	DB.Exec(database.DROP_TABLE)
	DB.Exec(database.CREATE_TABLE)

	t.Run("0 Records list 200OK but empty array -NO DATA-", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/expenses", nil)
		req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
		r.ServeHTTP(w, req)

		responseData, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, "[]", string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)
	})
	InsStmt, _ := DB.Prepare(database.INSERT)

	t.Log("Prepare 2 records")
	tags := []string{"beverage"}
	InsStmt.Exec("apple smoothie", 89, "no discount", pq.Array(&tags))
	tags = []string{"gadget"}
	InsStmt.Exec("iPhone 14 Pro Max 1TB", 66900, "birthday gift from my love", pq.Array(&tags))

	t.Run("List Expenses - 2 Records Return -SUCCESS-", func(t *testing.T) {
		w := httptest.NewRecorder()
		expectResponse := `[{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]},{"id":2,"title":"iPhone 14 Pro Max 1TB","amount":66900,"note":"birthday gift from my love","tags":["gadget"]}]`
		req, _ := http.NewRequest("GET", "/expenses", nil)
		req.SetBasicAuth(myusers.Good.Name, myusers.Good.Password)
		r.ServeHTTP(w, req)

		responseData, _ := ioutil.ReadAll(w.Body)
		assert.Equal(t, expectResponse, string(responseData))
		assert.Equal(t, http.StatusOK, w.Code)
	})

}

func TestBadAuthentication(t *testing.T) {
	t.Log("Test Case : Bad username")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/expenses", nil)
	req.SetBasicAuth(myusers.Bad.Name, myusers.Bad.Password)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

}
