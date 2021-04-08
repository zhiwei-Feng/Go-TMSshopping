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

func SelectProductCateById(id int) (domain.ProductCategory, error) {
	var target domain.ProductCategory
	result := db.DB.First(&target, id)
	if result.Error != nil {
		return target, result.Error
	}

	return target, nil
}

func SelectAllProductCate() ([]domain.ProductCategory, error) {
	var list []domain.ProductCategory
	result := db.DB.Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}
