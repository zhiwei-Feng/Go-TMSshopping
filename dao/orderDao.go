package dao

import (
	"strconv"
	"tmsshopping/db"
	"tmsshopping/domain"
)

func TotalPageForOrder(count int, id, name string) (int, error) {
	var (
		tpage = 1
		sum   int64
	)

	sql := "select count(*) from EASYBUY_ORDER where 1=1 "
	values := make([]interface{}, 0, 2)

	if id != "" {
		sql += " and EO_ID=?"
		values = append(values, id)
	}

	if name != "" {
		sql += " and EO_USER_ID like ? "
		values = append(values, "%"+name+"%")
	}

	result := db.DB.Debug().Raw(sql, values...).Scan(&sum)
	if result.Error != nil {
		return 0, result.Error
	}

	if sum%int64(count) == 0 {
		tpage = int(sum / int64(count))
	} else {
		tpage = int(sum/int64(count)) + 1
	}
	return tpage, nil
}

func SelectAllOrderForPagination(cpage, count int, id, name string) ([]domain.Order, error) {
	var list []domain.Order
	querySql := "select * from EASYBUY_ORDER where 1=1 "
	values := make([]interface{}, 0, 4)

	if id != "" {
		querySql += " and EO_ID=?"
		values = append(values, id)
	}

	if name != "" {
		querySql += " and EO_USER_ID like ? "
		values = append(values, "%"+name+"%")
	}

	querySql += " order by EO_ID desc limit ?,?"

	sql := "select * from(" +
		"select row_number() over () rn,a.* from(" +
		querySql +
		")a)b where b.rn between ? and ?"

	values = append(values, strconv.Itoa(count*(cpage-1)))
	values = append(values, strconv.Itoa(count))
	values = append(values, strconv.Itoa(count*(cpage-1)))
	values = append(values, strconv.Itoa(count))

	result := db.DB.Debug().Raw(sql, values...).Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}
