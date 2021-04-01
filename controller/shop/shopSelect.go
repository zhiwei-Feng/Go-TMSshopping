package shopController

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"tmsshopping/dao"
	"tmsshopping/domain"
)

// 购物车
func ShopSelect(ctx *gin.Context) {
	attributes := gin.H{}
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
		return
	}

	user, ok := loginUser.(domain.User)
	if ok {
		// 根据用户id获取其购物车列表
		if name, ok := sess.Get("name"); ok {
			attributes["name"] = name // 登录用户状态载入
		}
		shopList, _ := dao.GetShopCartOfUser(user.Id)
		attributes["list"] = shopList
		ctx.HTML(http.StatusOK, "shopping.tmpl", attributes)
	} else {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
}
