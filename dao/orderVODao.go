package dao

import (
	"tmsshopping/db"
	"tmsshopping/domain"
)

func SelectOrderVOByUsername(username string) ([]domain.OrderVO, error) {
	var orderList []domain.OrderVO
	query := "select * from EASYBUY_ORDER eo,EASYBUY_ORDER_detail eod,EASYBUY_PRODUCT ep where eo.eo_user_id=? and eod.eo_id=eo.eo_id and eod.ep_id= ep.ep_id order by eo.eo_id desc"
	results := db.DB.Raw(query, username).Scan(&orderList)
	if results.Error != nil {
		return nil, results.Error
	}

	return orderList, nil
}
