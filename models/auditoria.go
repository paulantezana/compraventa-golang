package models

import "github.com/jinzhu/gorm"

type Auditoria struct {
	gorm.Model
}

func (Auditoria) TableName() string {
	return "auditorias"
}