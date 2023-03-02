package models

import "github.com/dollarkillerx/eim/internal/generated"

type User struct {
	BasicModel
	Account  string         `gorm:"type:varchar(300);uniqueIndex" json:"account"`
	Name     string         `gorm:"type:varchar(300)" json:"name"`
	Password string         `gorm:"type:varchar(600)" json:"password"`
	Role     generated.Role `gorm:"type:varchar(60);index" json:"role"`
}
