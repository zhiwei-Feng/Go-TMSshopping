package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"tmsshopping/dao"
	"tmsshopping/db"
	"tmsshopping/domain"
	"tmsshopping/log"
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

func ProductAdd(ctx *gin.Context) {
	var (
		pnameStr = ctx.PostForm("productName")
		idStr    = ctx.PostForm("parentId")
		priceStr = ctx.PostForm("productPrice")
		descStr  = ctx.PostForm("productDesc")
		stockStr = ctx.PostForm("productStock")
		stock    int
		price    int
		err      error
	)

	log.Log.WithField("productName", pnameStr).
		WithField("parentId", idStr).
		WithField("productPrice", priceStr).
		Info("输入参数")

	// handle upload file
	photo, err := ctx.FormFile("photo")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	//dst := filepath.Base(photo.Filename)
	// 上传的文件存放在main.go同级目录的templates/imgs目录中
	dst := filepath.Join("images", "product", photo.Filename)
	log.Log.WithField("path", dst).Info("文件存储路径")
	if err := ctx.SaveUploadedFile(photo, dst); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	if stock, err = strconv.Atoi(stockStr); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if price, err = strconv.Atoi(priceStr); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	log.Log.WithField("idStr", idStr).Info("parentId:")
	ids := strings.Split(idStr, "-")
	if len(ids) != 2 {
		ctx.String(http.StatusBadRequest, "parentId 错误")
		return
	}
	cid, err1 := strconv.Atoi(ids[0])
	ccid, err2 := strconv.Atoi(ids[1])
	if err1 != nil || err2 != nil {
		ctx.String(http.StatusBadRequest, "parentId 错误")
		return
	}
	newProd := domain.Product{
		Name:            pnameStr,
		Description:     descStr,
		Price:           float32(price),
		Stock:           stock,
		CategoryId:      cid,
		CategoryChildId: ccid,
		FileName:        photo.Filename,
	}
	log.Log.WithField("newProd", newProd).Info("新增产品")

	result := db.DB.Create(&newProd)
	if result.Error != nil {
		ctx.String(http.StatusInternalServerError, "fail to add product.")
		return
	}

	ctx.Redirect(http.StatusFound, "productSelect")
}
