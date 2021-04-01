package shopController

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"tmsshopping/dao"
	"tmsshopping/domain"
)

// 商品购买按钮触发, GET method, 对应ShopAddServlet
func ShopBuy(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		paramPid   string
		paramCount string
		targetProd domain.Product // 购买的商品
	)

	paramPid = ctx.Query("id")
	paramCount = ctx.Query("count")
	if paramPid == "" {
		ctx.HTML(http.StatusOK, "shop_buy_err.html", gin.H{"error": "参数不正确"})
		return
	}
	pid, err := strconv.Atoi(paramPid)
	if err != nil {
		ctx.HTML(http.StatusOK, "shop_buy_err.html", gin.H{"error": "商品不存在"})
		return
	}
	count, err := strconv.Atoi(paramCount)
	if err != nil {
		ctx.HTML(http.StatusOK, "shop_buy_err.html", gin.H{"error": "参数不正确"})
		return
	}
	targetProd, _ = dao.SelectProductById(pid)
	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
		return
	}
	user, ok := loginUser.(domain.User)
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
		return
	}

	// 插入当前商品
	_, err = dao.AddToShopCart(targetProd, count, user.Id)
	if err != nil {
		ctx.HTML(http.StatusOK, "shop_buy_err.html", gin.H{"error": "购买失败"})
		return
	}
	ctx.Redirect(http.StatusFound, "/ShopSelect")
}
