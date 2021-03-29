package dao

import (
	"tmsshopping/db"
	"tmsshopping/domain"
)

// 根据用户名和密码查询用户个数，用于登录时判断是否有该用户, 对应原项目中的selectByNM
func CountUserByNM(name, pwd string) (int, error) {
	var count int64
	result := db.DB.Model(&domain.User{}).Where("EU_USER_ID = ? AND EU_PASSWORD = ?", name, pwd).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}

	return int(count), nil
}

// 抽象出来的根据用户名和密码查询用户的接口，替代原项目中相关的冗余代码
func SelectUserByNM(name, pwd string) (domain.User, error) {
	var target domain.User
	result := db.DB.Where("EU_USER_ID = ? AND EU_PASSWORD = ?", name, pwd).First(&target)
	if result.Error != nil {
		return target, result.Error
	}

	return target, nil
}
