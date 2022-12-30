package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/myserver"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func TestInitData(t *testing.T) {
	database.ConnectDB()
	database.DB.Exec(database.DROP_TABLE)
	database.DB.Exec(database.CREATETABLE)
	database.DB.Exec(database.TEST_RECORD1)
	database.DB.Exec(database.TEST_RECORD2)
	database.DB.Exec(database.TEST_RECORD3)
}

func TestRoot(t *testing.T) {
	expectResponse := `{"message":"Hello KTBG Go01"}`

	r = myserver.StartAndRoute()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoryExp02(t *testing.T) {
	expectResponse := `{"id":1,"title":"Title1","amount":1111.11,"note":"Note1","tags":["tags1","tags2"]}`

	req, _ := http.NewRequest("GET", "/expenses/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestStoryExp01(t *testing.T) {

	jsonRequest := `{"title":"Title4","amount":4444.44,"note":"Note4","tags":["tags2","tags3"]}`
	expectResponse := `{"id":4,"title":"Title4","amount":4444.44,"note":"Note4","tags":["tags2","tags3"]}`

	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer([]byte(jsonRequest)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
