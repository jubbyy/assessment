package model

import (
	"encoding/json"
	"fmt"
)

type Expense struct {
	Id     int64    `json:"id"`
	Title  string   `json:"title"`
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
		panic(err)
	}
	return fmt.Sprintf("%v", string(e))

}

func (expense Expense) FromJSON(rawstring string) Expense {
	var es Expense
	err := json.Unmarshal([]byte(rawstring), &es)
	if err != nil {
		panic(err)
	}
	return es
}

func NewExpenseFromJSON(jsonstring string) Expense {
	var es Expense
	fmt.Printf("%T", jsonstring)
	fmt.Printf("%T", []byte(jsonstring))

	err := json.Unmarshal([]byte(jsonstring), &es)
	if err != nil {
		panic(err)
	}

	return es
}
