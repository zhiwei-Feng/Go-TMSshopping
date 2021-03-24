package domain

type Product struct {
	Id              int     `gorm:"primaryKey;column:EP_ID"`
	Name            string  `gorm:"column:EP_NAME"`
	Description     string  `gorm:"column:EP_DESCRIPTION"`
	Price           float32 `gorm:"column:EP_PRICE"`
	Stock           int     `gorm:"column:EP_STOCK"`
	CategoryId      int     `gorm:"column:EPC_ID"`
	CategoryChildId int     `gorm:"column:EPC_CHILD_ID"`
	FileName        string  `gorm:"column:EP_FILE_NAME"`
}

func (Product) TableName() string {
	return "EASYBUY_PRODUCT"
}
