package model

type Customer struct {
	ID             uint32   `json:"id"`
	FirstName      string `json:"firstName"`
	SecondName     string `json:"secondName"`
	LastNameP      string `json:"lastNameP"`
	LastNameM      string `json:"lastNameM"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	Password       string `json:"passwordHash"`
	HashedPassword string
}
