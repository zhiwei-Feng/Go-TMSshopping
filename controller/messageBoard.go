package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"tmsshopping/dao"
	"tmsshopping/domain"
)

/*
对应原项目的SelallServlet(搞不懂这名字是什么鬼意思)
*/

func MessageBoard(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	var (
		page       = 1
		pagesize   = 2
		spage      string
		attributes = gin.H{}
		maxPage    int
		flist      []domain.ProductCategory
		clist      []domain.ProductCategory
		list       []domain.Comment
	)
	spage = ctx.Query("page")
	logrus.Info("get page:", spage)
	if spage != "" {
		if v, err := strconv.Atoi(spage); err == nil {
			page = v
		}
	}
	logrus.Info("current page:", page)

	flist, _ = dao.SelectProductCateFather()            //所有商品分类
	clist, _ = dao.SelectProductCateChild()             //所有子类别
	list, _ = dao.CommentPage(page, pagesize)           //当前页的评论
	maxPage, _ = dao.MaxCommentPageNum(int64(pagesize)) //评论最大页数

	// 用于页面遍历
	pageList := make([]int, maxPage)
	for i := 0; i < maxPage; i++ {
		pageList[i] = i + 1
	}

	attributes["list"] = list
	attributes["max_page"] = maxPage
	attributes["page"] = page
	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["pageList"] = pageList

	ctx.HTML(http.StatusOK, "guestbook.tmpl", attributes)
}
