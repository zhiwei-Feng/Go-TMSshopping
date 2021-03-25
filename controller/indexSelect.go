package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"tmsshopping/dao"
)

func IndexSelect(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather()
	clist, _ := dao.SelectProductCateChild()
	tlist, _ := dao.SelectProductsByT()
	hlist, _ := dao.SelectProductsByHot()

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
	attributes["search_words"] = session.Get("search_words")
	attributes["name"] = session.Get("name")

	ctx.HTML(http.StatusOK, "index.tmpl", attributes)
}
