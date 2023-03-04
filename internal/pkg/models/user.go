package models

type User struct {
	BasicModel
	Account  string `gorm:"type:varchar(300);uniqueIndex" json:"account"`
	Nickname string `gorm:"type:varchar(300)" json:"nickname"`
	Password string `gorm:"type:varchar(600)" json:"password"`
	Sex      bool   `json:"sex"` // 0 man: 1: female
}
