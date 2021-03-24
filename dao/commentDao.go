package dao

import (
	"tmsshopping/db"
	"tmsshopping/domain"
)

func SelectAllComments() ([]domain.Comment, error) {
	allComments := make([]domain.Comment, 0, 100)
	result := db.DB.Find(&allComments)
	if result.Error != nil {
		return nil, result.Error
	}
	return allComments, nil
}
