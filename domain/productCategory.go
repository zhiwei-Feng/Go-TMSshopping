package domain

type ProductCategory struct {
	Id       int    `gorm:"primaryKey;column:EPC_ID"`
	Name     string `gorm:"column:EPC_NAME"`
	ParentId int    `gorm:"column:EPC_PARENT_ID"`
}

func (ProductCategory) TableName() string {
	return "EASYBUY_PRODUCT_CATEGORY"
}
