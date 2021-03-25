package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"tmsshopping/dao"
)

func IndexSelect(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather() //所有商品分类
	clist, _ := dao.SelectProductCateChild()  //所有子类别
	tlist, _ := dao.SelectProductsByT()       //滚动区域和今日特价
	hlist, _ := dao.SelectProductsByHot()     //热卖

	session := sessions.Default(ctx)
	ids, ok := session.Get("ids").([]int)
	if ok {
		lastlyList, _ := dao.SelectProductsByIds(ids)
		attributes["lastlylist"] = lastlyList
	}

	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["tlist"] = tlist
	attributes["hlist"] = hlist
	attributes["search_words"] = session.Get("search_words") // 搜索框内容
	attributes["name"] = session.Get("name")                 //登录用户

	ctx.HTML(http.StatusOK, "index.tmpl", attributes)
}
