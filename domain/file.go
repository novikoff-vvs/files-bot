package domain

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string
}

func NewFile(id int, name string) File {
	return File{
		Model: gorm.Model{ID: uint(id)},
		Name:  name,
	}
}
