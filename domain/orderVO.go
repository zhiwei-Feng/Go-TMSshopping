package domain

type OrderVO struct {
	Id           int     `gorm:"column:EO_ID"`
	UserId       string  `gorm:"column:EO_USER_ID"`
	ProdName     string  `gorm:"column:EP_NAME"`
	ProdFileName string  `gorm:"column:EP_FILE_NAME"`
	ProdPrice    float32 `gorm:"column:EP_PRICE"`
	Quantity     int     `gorm:"column:EOD_QUANTITY"`
	ProdStock    int     `gorm:"column:EP_STOCK"`
}
