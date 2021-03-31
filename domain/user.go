package domain

import "time"

type User struct {
	Id       string    `gorm:"primaryKey;column:EU_USER_ID;not null"`
	UserName string    `gorm:"column:EU_USER_NAME"`
	Password string    `gorm:"column:EU_PASSWORD"`
	Sex      string    `gorm:"column:EU_SEX"`
	Birthday time.Time `gorm:"column:EU_BIRTHDAY"`
	Email    string    `gorm:"column:EU_EMAIL"`
	Mobile   string    `gorm:"column:EU_MOBILE"`
	Address  string    `gorm:"column:EU_ADDRESS"`
	Status   int       `gorm:"column:EU_STATUS"`
}

func (User) TableName() string {
	return "EASYBUY_USER"
}
