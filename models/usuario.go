package models

import "github.com/jinzhu/gorm"

type Usuario struct {
	gorm.Model
}

func (Usuario) TableName() string {
	return "usuarios"
}