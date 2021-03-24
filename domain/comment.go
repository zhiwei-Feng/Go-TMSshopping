package domain

type Comment struct {
	Id         int    `gorm:"primaryKey;column:EC_ID"`
	Content    string `gorm:"column:EC_CONTENT"`
	CreateTime string `gorm:"column:EC_CREATE_TIME"`
	Reply      string `gorm:"column:EC_REPLY"`
	ReplyTime  string `gorm:"column:EC_REPLY_TIME"`
	NickName   string `gorm:"column:EC_NICK_NAME"`
}

func (Comment) TableName() string {
	return "EASYBUY_COMMENT"
}
