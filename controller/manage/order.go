package manage

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"tmsshopping/dao"
	"tmsshopping/db"
	"tmsshopping/domain"
)

// manage/ordersel: OrderSelServlet
func OrderManagePage(ctx *gin.Context) {
	var (
		cpage      = 1
		count      = 10
		cp         = ctx.Query("cp")
		idStr      = ctx.Query("orderId")
		name       = ctx.Query("userName")
		tpage      = 0
		orders     []domain.Order
		attributes = gin.H{}
	)

	if v, err := strconv.Atoi(cp); err == nil {
		cpage = v
	}

	tpage, _ = dao.TotalPageForOrder(count, idStr, name)
	orders, _ = dao.SelectAllOrderForPagination(cpage, count, idStr, name)

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

	attributes["order"] = orders
	attributes["cpage"] = cpage
	attributes["tpage"] = tpage
	attributes["orderId"] = idStr
	attributes["userName"] = name
	attributes["name"] = user
	attributes["plist"] = plist

	ctx.HTML(http.StatusOK, "order-manage.tmpl", attributes)
}

func OrderDelete(ctx *gin.Context) {
	var (
		idStr = ctx.Query("id")
		id    int
		err   error
	)

	if id, err = strconv.Atoi(idStr); err != nil {
		ctx.HTML(http.StatusBadRequest, "order_del_err.html", gin.H{})
		return
	}

	// todo: 目前按照原项目的逻辑来的，但是此处存在order-detail表中会遗留大量无用数据
	result := db.DB.Delete(&domain.Order{}, id)
	if result.Error != nil {
		ctx.Header("Content-Type", "text/html;charset=utf-8")
		ctx.HTML(http.StatusInternalServerError, "order_del_err.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
}

func OrderUpdatePage(ctx *gin.Context) {
	var (
		idStr      = ctx.Query("id")
		id         int
		order      domain.Order
		dlist      []domain.OrderVO
		attributes = gin.H{}
	)

	id, _ = strconv.Atoi(idStr)

	_ = db.DB.First(&order, id)
	dlist, _ = dao.SelectOrderVOById(id)

	sess := ginsession.FromContext(ctx)
	if name, ok := sess.Get("name"); ok {
		attributes["name"] = name // 登录用户状态载入
	}

	attributes["dlist"] = dlist
	attributes["order"] = order
	attributes["statuslist"] = domain.StatusList

	ctx.HTML(http.StatusOK, "order-modify.tmpl", attributes)
}

func OrderUpdate(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		paramOrderId = ctx.Query("orderId")
		paramStatus  = ctx.Query("tatus")
		orderId      int
		status       int
		err          error
	)

	if orderId, err = strconv.Atoi(paramOrderId); err != nil {
		ctx.HTML(http.StatusBadRequest, "order_mod_err.html", gin.H{})
		return
	}

	if status, err = strconv.Atoi(paramStatus); err != nil {
		ctx.HTML(http.StatusBadRequest, "order_mod_err.html", gin.H{})
		return
	}

	upOrder := domain.Order{
		Id:     orderId,
		Status: status,
	}

	result := db.DB.Model(&upOrder).Select("EO_STATUS").Updates(&upOrder)
	if result.Error != nil {
		ctx.HTML(http.StatusInternalServerError, "order_mod_err.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
}
