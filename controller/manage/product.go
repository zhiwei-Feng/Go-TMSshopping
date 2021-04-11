package manage

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"tmsshopping/dao"
	"tmsshopping/domain"
)

func ProductManagePage(ctx *gin.Context) {
	var (
		cpage      = 1
		count      = 5
		tpage      = 0
		cp         = ctx.Query("cp")
		attributes = gin.H{}
	)

	if v, err := strconv.Atoi(cp); err == nil {
		cpage = v
	}

	eplist, _ := dao.SelectAllProducts(cpage, count)
	if v, err := dao.TotalPageOfProducts(count); err == nil {
		tpage = v
	}

	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	user, ok := loginUser.(domain.User)
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}

	// 给前端跳转页面用的变量
	plist := make([]int, tpage)
	for i := 1; i <= tpage; i++ {
		plist[i-1] = i
	}

	attributes["eplist"] = eplist
	attributes["cpage"] = cpage
	attributes["tpage"] = tpage
	attributes["name"] = user
	attributes["plist"] = plist

	ctx.HTML(http.StatusOK, "product-manage.tmpl", attributes)
}

func ProductAddPage(ctx *gin.Context) {
	var (
		attributes = gin.H{}
		flist      []domain.ProductCategory
		clist      []domain.ProductCategory
	)

	flist, _ = dao.SelectProductCateFather()
	clist, _ = dao.SelectProductCateChild()

	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	user, ok := loginUser.(domain.User)
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}

	attributes["flist"] = flist
	attributes["clist"] = clist
	attributes["name"] = user

	ctx.HTML(http.StatusOK, "product-add.tmpl", attributes)
}
