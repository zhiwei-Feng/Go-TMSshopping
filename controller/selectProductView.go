package controller

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"tmsshopping/dao"
)

func SelectProductView(ctx *gin.Context) {
	attributes := gin.H{}
	flist, _ := dao.SelectProductCateFather()
	clist, _ := dao.SelectProductCateChild()
	id := ctx.Query("id")
	session := ginsession.FromContext(ctx)
	idsStr, ok := session.Get("ids")
	var ids []int
	if !ok {
		ids = []int{}
	} else {
		ids = idsStr.([]int)
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
	session.Save()

	// set attributes
	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["lastlylist"] = lastlyList

	user, _ := session.Get("name")
	attributes["name"] = user

	ctx.HTML(http.StatusOK, "product-view.tmpl", attributes)
}
