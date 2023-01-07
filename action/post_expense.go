package action

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jubbyy/assessment/database"
	"github.com/jubbyy/assessment/model"
	"github.com/lib/pq"
)

func PostExpense2(c *gin.Context) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format JSON data (action.PostExpense)"})
		return
	}

	//	tags := strings.Join(e.Tags, ",")
	err := database.PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&e.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Insertion to database error"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func PostExpense(c *gin.Context, dh *database.DatabaseHandler) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format JSON data (action.PostExpense)"})
		return
	}
	//	c.JSON(http.StatusOK, e)
	//	err := dh.PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&e.Id)
	//	code, eo := dh.Create(e)
	code, eo := dh.Create(e)
	/*	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Insertion to database error"})
		return
	}*/
	c.JSON(code, eo)
}
func CreateExpense(c *gin.Context, db *sql.DB) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format JSON data (action.PostExpense)"})
		return
	}
	//	c.JSON(http.StatusOK, e)
	//	err := dh.PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&e.Id)
	//	code, eo := dh.Create(e)
	PostStmt, err := db.Prepare(database.INSERT)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on prepare statement"})
	}
	defer PostStmt.Close()

	er := PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&e.Id)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Insertion to database error"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func AddExpense(c *gin.Context, sg *database.StatementGroup) {
	var e model.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format JSON data (action.PostExpense)"})
		return
	}
	er := sg.PostStmt.QueryRow(e.Title, e.Amount, e.Note, pq.Array(&e.Tags)).Scan(&e.Id)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Insertion to database error"})
		return
	}
	c.JSON(http.StatusOK, e)
}
