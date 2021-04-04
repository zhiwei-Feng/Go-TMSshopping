package dao

import (
	"gorm.io/gorm"
	"time"
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

func DeleteItem(id int) (int, error) {
	result := db.DB.Delete(&domain.ShopCart{}, id)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

// ============= 结算相关dao方法 =============
// 结算时更新产品库存, 涉及到事务需要传入对应的事务*gorm.DB
func UpdateStock(id, stock int, DBCon *gorm.DB) (int, error) {
	var prod = domain.Product{Id: id}
	result := DBCon.Model(&prod).UpdateColumn("EP_STOCK", gorm.Expr("EP_STOCK - ?", stock))
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func CreateOrder(id, name, address string, price int, DBCon *gorm.DB) (int, error) {
	newOrder := domain.Order{
		UserId:      id,
		Username:    name,
		UserAddress: address,
		CreateTime:  time.Now(),
		Cost:        float32(price),
		Status:      1, // hard code
		Type:        1, // hard code
	}
	result := DBCon.Create(&newOrder)
	if result.Error != nil {
		return 0, result.Error
	}

	return newOrder.Id, nil
}

func GenerateOrderDetail(orderId, prodId, quan, cost int, DBCon *gorm.DB) (int, error) {
	detail := domain.OrderDetail{
		OrderId:   orderId,
		ProductId: prodId,
		Quantity:  quan,
		Cost:      float32(cost),
	}
	result := DBCon.Create(&detail)
	if result.Error != nil {
		return 0, result.Error
	}

	return detail.Id, nil
}

// 结算时，更新购物车条目状态
func SettleItem(id int, DBCon *gorm.DB) (int, error) {
	var item = domain.ShopCart{Id: id}
	result := DBCon.Model(&item).Update("Valid", 2)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

// =============================================================
