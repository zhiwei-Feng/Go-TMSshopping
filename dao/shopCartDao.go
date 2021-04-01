package dao

import (
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
