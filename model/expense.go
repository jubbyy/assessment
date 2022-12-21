package model

type Expenses struct {
	Id     int64
	Title  string
	Amount int
	Note   string
	Tags   []string
}
