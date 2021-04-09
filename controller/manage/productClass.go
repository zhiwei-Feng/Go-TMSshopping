package manage

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"tmsshopping/dao"
	"tmsshopping/db"
	"tmsshopping/domain"
	"tmsshopping/log"
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

func ProductClassAddPage(ctx *gin.Context) {
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

	ctx.HTML(http.StatusOK, "productClass-add.tmpl", attributes)
}

func ProductClassAdd(ctx *gin.Context) {
	var (
		fid    int
		fidStr = ctx.Query("parentId")
		name   = ctx.Query("className")
		err    error
	)

	if fid, err = strconv.Atoi(fidStr); err != nil {
		log.Log.WithField("err", err).Error("参数转换出错")
		ctx.HTML(http.StatusInternalServerError, "productClass_add_err.html", gin.H{})
		return
	}

	newCate := domain.ProductCategory{
		Name:     name,
		ParentId: fid,
	}
	result := db.DB.Create(&newCate)
	if result.Error != nil {
		ctx.HTML(http.StatusInternalServerError, "productClass_add_err.html", gin.H{})
		return
	}

	ctx.Redirect(http.StatusFound, "productClass")
}

func ProductClassDel(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		result := db.DB.Delete(&domain.ProductCategory{}, id)
		if result.Error != nil {
			ctx.HTML(http.StatusInternalServerError, "productClass_del_err.html", gin.H{})
			return
		}

		ctx.Redirect(http.StatusFound, "productClass")
		return
	}

	ctx.HTML(http.StatusInternalServerError, "productClass_del_err.html", gin.H{})
}

func ProductClassUpdatePage(ctx *gin.Context) {
	var (
		idStr      = ctx.Query("id")
		attributes = gin.H{}
		id         int
		err        error
	)

	if id, err = strconv.Atoi(idStr); err != nil {
		ctx.HTML(http.StatusBadRequest, "productClass_mod_err.html", gin.H{})
		return
	}

	epc, err := dao.SelectProductCateById(id)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "productClass_mod_err.html", gin.H{})
		return
	}

	epclist, err := dao.SelectAllProductCate()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "productClass_mod_err.html", gin.H{})
		return
	}

	attributes["epc"] = epc
	attributes["epclist"] = epclist

	ctx.HTML(http.StatusOK, "productClass-modify.tmpl", attributes)
}
