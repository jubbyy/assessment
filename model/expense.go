package model

import (
	"encoding/json"
	"fmt"
)

type Expense struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float32  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func (expense Expense) AsJSON() string {
	e, err := json.Marshal(expense)
	if err != nil {
		panic(e)
	}
	return fmt.Sprintf("%v", string(e))
}
