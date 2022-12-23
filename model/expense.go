package model

type Expenses struct {
	Id     int64
	Title  string
	Amount float32
	Note   string
	Tags   []string
}
