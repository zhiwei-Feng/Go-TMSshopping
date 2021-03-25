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
		"(select ep_id eod_ep_id,sum(EOD_QUANTITY) buysum from EASYBUY_ORDER_detail " +
		"group by EP_id order by sum(EOD_QUANTITY) desc) b  " +
		"where a.ep_id=b.eod_ep_id order by buysum desc  ) tab1) tab2 limit 0,8"

	result := db.DB.Raw(sql).Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

func TotalPageOfProducts(count int) (int, error) {
	var (
		tpage      int64
		totalCount int64
	)
	tpage = 1
	results := db.DB.Model(&domain.Product{}).Count(&totalCount)
	if results.Error != nil {
		return -1, results.Error
	}

	if totalCount%int64(count) == 0 {
		tpage = totalCount / int64(count)
	} else {
		tpage = totalCount/int64(count) + 1
	}
	return int(tpage), nil
}

func TotalPageOfProductsByFid(count, fid int) (int, error) {
	var (
		tpage      int64
		totalCount int64
	)
	tpage = 1
	results := db.DB.Model(&domain.Product{}).Where("EPC_ID = ?", fid).Count(&totalCount)
	if results.Error != nil {
		return -1, results.Error
	}

	if totalCount%int64(count) == 0 {
		tpage = totalCount / int64(count)
	} else {
		tpage = totalCount/int64(count) + 1
	}
	return int(tpage), nil
}

func TotalPageOfProductsByCid(count, cid int) (int, error) {
	var (
		tpage      int64
		totalCount int64
	)
	tpage = 1
	results := db.DB.Model(&domain.Product{}).Where("EPC_CHILD_ID = ?", cid).Count(&totalCount)
	if results.Error != nil {
		return -1, results.Error
	}

	if totalCount%int64(count) == 0 {
		tpage = totalCount / int64(count)
	} else {
		tpage = totalCount/int64(count) + 1
	}
	return int(tpage), nil
}

func TotalPageOfProductsByName(count int, name string) (int, error) {
	var (
		tpage      int64
		totalCount int64
	)
	tpage = 1
	results := db.DB.Model(&domain.Product{}).Where("EP_NAME LIKE ?", "%"+name+"%").Count(&totalCount)
	if results.Error != nil {
		return -1, results.Error
	}

	if totalCount%int64(count) == 0 {
		tpage = totalCount / int64(count)
	} else {
		tpage = totalCount/int64(count) + 1
	}
	return int(tpage), nil
}

func SelectAllProductsByFid(cpage, count, fid int) ([]domain.Product, error) {
	var (
		productsOfF []domain.Product
	)

	results := db.DB.Where("EPC_ID = ?", fid).Order("EP_ID desc").Offset(count * (cpage - 1)).Limit(count).Find(&productsOfF)
	if results.Error != nil {
		return nil, results.Error
	}

	return productsOfF, nil
}

func SelectAllProducts(cpage, count int) ([]domain.Product, error) {
	var allProducts []domain.Product
	result := db.DB.Order("EP_ID desc").Offset(count * (cpage - 1)).Limit(count).Find(&allProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	return allProducts, nil
}

func SelectAllProductsByCid(cpage, count, cid int) ([]domain.Product, error) {
	var (
		productsOfC []domain.Product
	)

	results := db.DB.Where("EPC_CHILD_ID = ?", cid).Order("EP_ID desc").Offset(count * (cpage - 1)).Limit(count).Find(&productsOfC)
	if results.Error != nil {
		return nil, results.Error
	}

	return productsOfC, nil
}

func SelectAllProductsByName(name string) ([]domain.Product, error) {
	var (
		productsOfName []domain.Product
	)

	results := db.DB.Where("EP_NAME LIKE ?", "%"+name+"%").Find(&productsOfName)
	if results.Error != nil {
		return nil, results.Error
	}

	return productsOfName, nil
}
