package models

import "github.com/jinzhu/gorm"

type Cliente struct {
	gorm.Model
}

func (Cliente) TableName() string {
	return "clientes"
}