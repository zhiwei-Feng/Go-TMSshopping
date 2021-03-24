package dao

import (
	"tmsshopping/db"
	"tmsshopping/domain"
)

func SelectProductById(id int) (domain.Product, error) {
	target := domain.Product{}
	result := db.DB.First(&target, id)
	if result.Error != nil {
		return target, result.Error
	}

	return target, nil
}

func SelectProductsByIds(ids []int) ([]domain.Product, error) {
	var lastlyList []domain.Product
	result := db.DB.Where("EP_ID IN ?", ids).Find(&lastlyList)
	if result.Error != nil {
		return nil, result.Error
	}

	return lastlyList, nil
}

func SelectProductsByT() ([]domain.Product, error) {
	var list []domain.Product
	result := db.DB.Order("EP_PRICE asc").Offset(0).Limit(9).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

func SelectProductsByHot() ([]domain.Product, error) {
	var list []domain.Product
	sql := "select * from ( " +
		"select tab1.* from  (  " +
		"select * from EASYBUY_PRODUCT a,  " +
		"(select ep_id eod_ep_id,sum(EOD_QUANTITY) buysum from EASYBUY_ORDER_DETAIL " +
		"group by EP_id order by sum(EOD_QUANTITY) desc) b  " +
		"where a.ep_id=b.eod_ep_id order by buysum desc  ) tab1) tab2 limit 0,8"

	result := db.DB.Raw(sql).Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}
