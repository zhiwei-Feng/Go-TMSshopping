package domain

import "time"

type Comment struct {
	Id         int       `gorm:"primaryKey;column:EC_ID"`
	Content    string    `gorm:"column:EC_CONTENT" form:"guestContent"`
	CreateTime time.Time `gorm:"column:EC_CREATE_TIME"`
	Reply      string    `gorm:"column:EC_REPLY"`
	ReplyTime  time.Time `gorm:"column:EC_REPLY_TIME"`
	NickName   string    `gorm:"column:EC_NICK_NAME" form:"guestName"`
}

func (Comment) TableName() string {
	return "EASYBUY_COMMENT"
}
