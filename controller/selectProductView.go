package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tmsshopping/dao"
)

func SelectProductView(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather()
	clist, _ := dao.SelectProductCateChild()
	id := ctx.Query("id")
	session := sessions.Default(ctx)
	ids, ok := session.Get("ids").([]int)
	if !ok || ids == nil {
		ids = []int{}
	}
	if len(ids) >= 5 {
		ids = ids[1:]
	}

	if idConvert, err := strconv.Atoi(id); err == nil {
		ids = append(ids, idConvert)
		p, _ := dao.SelectProductById(idConvert)
		attributes["p"] = p
	}
	lastlyList, _ := dao.SelectProductsByIds(ids)
	session.Set("ids", ids)

	// set attributes
	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["lastlylist"] = lastlyList

	ctx.HTML(http.StatusOK, "product-view.tmpl", attributes)
}
