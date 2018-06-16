package models

import "github.com/jinzhu/gorm"

type Compra struct {
	gorm.Model
}

func (Compra) TableName() string {
	return "compras"
}