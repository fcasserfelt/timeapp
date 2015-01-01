package domain

type User struct {
	Id    int
	Email string
}

type UserRepository interface {
	FindByEmail(email string) User
}
