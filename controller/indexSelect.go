package controller

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"tmsshopping/dao"
)

func IndexSelect(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather() //所有商品分类
	clist, _ := dao.SelectProductCateChild()  //所有子类别
	tlist, _ := dao.SelectProductsByT()       //滚动区域和今日特价
	hlist, _ := dao.SelectProductsByHot()     //热卖

	sess := ginsession.FromContext(ctx)
	words, _ := sess.Get("search_words")
	user, _ := sess.Get("name")
	idsStr, ok := sess.Get("ids")
	if ok {
		ids := idsStr.([]int)
		lastlyList, _ := dao.SelectProductsByIds(ids)
		attributes["lastlylist"] = lastlyList
	}

	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["tlist"] = tlist
	attributes["hlist"] = hlist
	attributes["search_words"] = words // 搜索框内容
	attributes["name"] = user          //登录用户

	ctx.HTML(http.StatusOK, "index.tmpl", attributes)
}
