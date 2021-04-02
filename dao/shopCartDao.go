package dao

import (
	"gorm.io/gorm"
	"tmsshopping/db"
	"tmsshopping/domain"
)

func GetShopCartOfUser(username string) ([]domain.ShopCart, error) {
	var (
		shopList []domain.ShopCart
	)

	results := db.DB.Where("es_eu_user_id = ? AND es_valid = ?", username, 1).Order("es_id desc").Find(&shopList)
	if results.Error != nil {
		return nil, results.Error
	}

	return shopList, nil
}

func AddToShopCart(prod domain.Product, count int, username string) (int, error) {
	newShopCartItem := domain.ShopCart{
		ProductFileName: prod.FileName,
		ProductName:     prod.Name,
		ProductPrice:    prod.Price,
		Quantity:        count,
		Stock:           prod.Stock,
		ProductId:       prod.Id,
		UserId:          username,
		Valid:           1, // 硬编码, 不知道为啥
	}
	result := db.DB.Create(&newShopCartItem)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

// 购物车购买数量+1
func PlusItem(id int) (int, error) {
	var item = domain.ShopCart{Id: id}
	result := db.DB.Model(&item).UpdateColumn("es_eod_quantity", gorm.Expr("es_eod_quantity + ?", 1))
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func ReduceItem(id int) (int, error) {
	var item = domain.ShopCart{Id: id}
	result := db.DB.Model(&item).UpdateColumn("es_eod_quantity", gorm.Expr("es_eod_quantity - ?", 1))
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func SetItem(id, num int) (int, error) {
	var item = domain.ShopCart{Id: id}
	result := db.DB.Model(&item).Update("es_eod_quantity", num)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
