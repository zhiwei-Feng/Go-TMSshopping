package domain

type OrderDetail struct {
	Id        int     `gorm:"primaryKey;column:EOD_ID"`
	OrderId   int     `gorm:"column:EO_ID"`
	ProductId int     `gorm:"column:EP_ID"`
	Quantity  int     `gorm:"column:EOD_QUANTITY"`
	Cost      float32 `gorm:"column:EOD_COST"`
}

func (OrderDetail) TableName() string {
	return "EASYBUY_ORDER_detail"
}
