package entity

type UserEntity struct {
	ID         int
	Email      string
	Name       string
	Password   string
	IsVerified bool
	RoleName   string
	Address    string
	Phone      string
	Photo      string
	Lat        string
	Lng        string
}