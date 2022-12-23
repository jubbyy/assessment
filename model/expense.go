package model

import (
	"encoding/json"
	"fmt"
)

type Expense struct {
	Id     int64    `json:"id"`
	Title  string   `json:"tilte"`
	Amount float32  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Expenses []Expense

func (expense Expense) AsJSON() string {
	e, err := json.Marshal(expense)
	if err != nil {
		panic(e)
	}
	return fmt.Sprintf("%v", string(e))
}

func (expenses Expenses) AsJSON() string {
	e, err := json.Marshal(expenses)
	if err != nil {
		panic(e)
	}
	return fmt.Sprintf("%v", string(e))

}
