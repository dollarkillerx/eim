package models

type User struct {
	BasicModel        // 继承 BasicModel 结构体
	Account    string `gorm:"type:varchar(300);uniqueIndex" json:"account"` // 账号，长度300，设定唯一索引
	FullName   string `gorm:"type:varchar(300)" json:"full_name"`           // 姓名，长度300
	Nickname   string `gorm:"type:varchar(300)" json:"nickname"`            // 昵称，长度300
	Birthday   string `gorm:"type:varchar(300)" json:"birthday"`            // 生日，长度300
	Email      string `gorm:"type:varchar(300)" json:"email"`               // 邮箱，长度300
	About      string `gorm:"type:varchar(600)" json:"about"`               // 用户简介，长度600
	Password   string `gorm:"type:varchar(300)" json:"password"`            // 用户密码，长度300
	Avatar     string `gorm:"type:varchar(300)" json:"avatar"`
}
