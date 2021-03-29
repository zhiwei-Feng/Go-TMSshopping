package controller

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"tmsshopping/dao"
	"tmsshopping/domain"
)

func SelectProductList(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather() //所有商品分类
	clist, _ := dao.SelectProductCateChild()  //所有子类别
	session := ginsession.FromContext(ctx)
	idsStr, ok := session.Get("ids")
	if ok {
		ids := idsStr.([]int)
		lastlylist, _ := dao.SelectProductsByIds(ids)
		attributes["lastlylist"] = lastlylist
	}

	var (
		cpage int
		count int
		tpage int
		list  []domain.Product
	)

	cpage = 1
	count = 8
	tpage = 0
	cp := ctx.Query("cp")
	if cp != "" {
		cpage, _ = strconv.Atoi(cp)
	}

	fid := ctx.Query("fid")
	cid := ctx.Query("cid")
	name := ctx.Query("name")
	if fid == "" && cid == "" {
		list, _ = dao.SelectAllProducts(cpage, count)
		tpage, _ = dao.TotalPageOfProducts(count)
		attributes["title"] = "All Goods"
	}
	if fid != "" {
		id, err := strconv.Atoi(fid)
		if err == nil {
			list, _ = dao.SelectAllProductsByFid(cpage, count, id)
			tpage, _ = dao.TotalPageOfProductsByFid(count, id)
			productCateOfId, _ := dao.SelectProductCateById(id)
			attributes["title"] = productCateOfId.Name
		}
	}
	if cid != "" {
		id, err := strconv.Atoi(cid)
		if err == nil {
			list, _ = dao.SelectAllProductsByCid(cpage, count, id)
			tpage, _ = dao.TotalPageOfProductsByCid(count, id)
			productCateOfId, _ := dao.SelectProductCateById(id)
			attributes["title"] = productCateOfId.Name
		}
	}
	if name != "" {
		list, _ = dao.SelectAllProductsByName(name)
		tpage, _ = dao.TotalPageOfProductsByName(count, name)
	}

	// addition 构造数据为页面select所用
	selectList := make([]int, 0, tpage)
	for i := 1; i <= tpage; i++ {
		selectList = append(selectList, i)
	}

	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["list"] = list
	attributes["cpage"] = cpage
	attributes["tpage"] = tpage
	attributes["search_words"] = name
	attributes["selected_fid"] = fid
	attributes["selectList"] = selectList

	user, _ := session.Get("name")
	attributes["name"] = user

	ctx.HTML(http.StatusOK, "product-list.tmpl", attributes)
}
