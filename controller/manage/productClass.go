package manage

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"tmsshopping/dao"
	"tmsshopping/domain"
)

func ProductClassManagePage(ctx *gin.Context) {
	var (
		attributes = gin.H{}
		epclist    []domain.ProductCategory
	)

	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	user, ok := loginUser.(domain.User)
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}

	epclist, _ = dao.SelectAllProductCate()

	attributes["name"] = user
	attributes["epclist"] = epclist

	ctx.HTML(http.StatusOK, "productClass-manage.tmpl", attributes)
}
