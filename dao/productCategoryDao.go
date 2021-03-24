package dao

import (
	"tmsshopping/db"
	"tmsshopping/domain"
)

func SelectProductCateFather() ([]domain.ProductCategory, error) {
	var fathers []domain.ProductCategory
	result := db.DB.Where("EPC_PARENT_ID", 0).Find(&fathers)
	if result.Error != nil {
		return nil, result.Error
	}

	return fathers, nil
}

func SelectProductCateChild() ([]domain.ProductCategory, error) {
	var childs []domain.ProductCategory
	result := db.DB.Where("EPC_ID!=EPC_PARENT_ID").Find(&childs)
	if result.Error != nil {
		return nil, result.Error
	}

	return childs, nil
}
