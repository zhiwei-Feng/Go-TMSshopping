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

// 对应selPage
func CommentPage(page, pagesize int) ([]domain.Comment, error) {
	var (
		min         = (page - 1) * pagesize
		max         = pagesize
		allComments []domain.Comment
	)

	results := db.DB.Order("EC_CREATE_TIME desc").Offset(min).Limit(max).Find(&allComments)
	if results.Error != nil {
		return nil, results.Error
	}

	return allComments, nil
}

// getMax
func MaxCommentPageNum(pagesize int64) (int, error) {
	var (
		max   int64
		count int64
	)

	result := db.DB.Model(&domain.Comment{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}

	if count%pagesize == 0 {
		max = count / pagesize
	} else {
		max = (count / pagesize) + 1
	}
	return int(max), nil
}
