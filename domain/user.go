package domain

import "gorm.io/gorm"

const WAIT_INPUT = "wait_input"
const GET_FILE = "get_file"

type User struct {
	gorm.Model
	Terminator string
}

func NewUser(ID int64) User {
	return User{
		Model: gorm.Model{ID: uint(ID)},
	}
}
