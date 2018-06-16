package models

import "github.com/jinzhu/gorm"

type Producto struct {
	gorm.Model
}

func (Producto) TableName() string {
	return "productos"
}