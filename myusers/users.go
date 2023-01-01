package myusers

type User struct {
	Name     string
	Password string
}

var (
	Good  = User{"user", "1234"}
	Admin = User{"user", "1234"}
	Bad   = User{"user", "1221"}
)
