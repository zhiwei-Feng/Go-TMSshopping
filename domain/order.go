package domain

import "time"

type Order struct {
	Id          int       `gorm:"primaryKey;column:EO_ID"`
	UserId      string    `gorm:"column:EO_USER_ID"`
	Username    string    `gorm:"column:EO_USER_NAME"`
	UserAddress string    `gorm:"column:EO_USER_ADDRESS"`
	CreateTime  time.Time `gorm:"column:EO_CREATE_TIME"`
	Cost        float32   `gorm:"column:EO_COST"`
	Status      int       `gorm:"column:EO_STATUS"`
	Type        int       `gorm:"column:EO_TYPE"`
}

func (Order) TableName() string {
	return "EASYBUY_ORDER"
}
