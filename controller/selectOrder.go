package controller

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"tmsshopping/dao"
)

// 个人订单 -> SelectDD
func SelectOrder(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather() //所有商品分类
	clist, _ := dao.SelectProductCateChild()  //所有子类别
	tlist, _ := dao.SelectProductsByT()       //滚动区域和今日特价
	hlist, _ := dao.SelectProductsByHot()     //热卖

	sess := ginsession.FromContext(ctx)
	idsStr, ok := sess.Get("ids")
	if ok {
		ids := idsStr.([]int)
		lastlyList, _ := dao.SelectProductsByIds(ids)
		attributes["lastlylist"] = lastlyList
	}
	user, _ := sess.Get("name")

	dd := ctx.Query("dd")
	dList, _ := dao.SelectOrderVOByUsername(dd)
	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["tlist"] = tlist
	attributes["hlist"] = hlist
	attributes["dan"] = dList
	attributes["name"] = user

	ctx.HTML(http.StatusOK, "order.tmpl", attributes)
}
