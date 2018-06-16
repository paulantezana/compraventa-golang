package models

import "github.com/jinzhu/gorm"

type Perfil struct {
	gorm.Model
}

func (Perfil) TableName() string {
	return "perfiles"
}