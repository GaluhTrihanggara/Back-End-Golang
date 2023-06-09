package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Id       uint   `json:"id" form:"id"`
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}
