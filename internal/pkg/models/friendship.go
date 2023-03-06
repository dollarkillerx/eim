package models

// Friendship ...
type Friendship struct {
	BasicModel        // 继承 BasicModel 结构体
	User1ID    string `gorm:"not null;index" json:"user1_id"` // 用户1在关系中的ID，非空，设定索引
	User2ID    string `gorm:"not null;index" json:"user2_id"` // 用户2在关系中的ID，非空，设定索引
}
