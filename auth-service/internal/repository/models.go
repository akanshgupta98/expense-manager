package repository

type Models struct {
	User User
}

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
}
