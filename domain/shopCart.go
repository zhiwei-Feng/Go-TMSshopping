package domain

// 购物车实体
type ShopCart struct {
	Id              int     `gorm:"primaryKey;column:es_id"`
	ProductFileName string  `gorm:"column:es_ep_file_name"`
	ProductName     string  `gorm:"column:es_ep_name"`
	ProductPrice    float32 `gorm:"column:es_ep_price"`
	Quantity        int     `gorm:"column:es_eod_quantity"`
	Stock           int     `gorm:"column:es_ep_stock"`
	ProductId       int     `gorm:"column:es_ep_id"`
	UserId          string  `gorm:"column:es_eu_user_id"`
	Valid           int     `gorm:"column:es_valid"`
}

func (ShopCart) TableName() string {
	return "EASYBUY_SHOP"
}
