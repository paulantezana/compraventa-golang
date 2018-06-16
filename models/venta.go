package models

import "github.com/jinzhu/gorm"

type Venta struct {
	gorm.Model
}

func (Venta) TableName() string {
	return "ventas"
}