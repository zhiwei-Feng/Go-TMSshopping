package shopController

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"tmsshopping/dao"
	"tmsshopping/db"
	"tmsshopping/domain"
)

func ShopCartSettle(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		sess = ginsession.FromContext(ctx)
		// params
		price    string
		epId     []string
		quantity []string
		spPrice  []string
		sid      []string
		esId     []string
	)

	// 获取当前登录用户
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

	// 获取参数
	price = ctx.Query("jstext")
	epId = ctx.QueryArray("spID")
	quantity = ctx.QueryArray("number")
	spPrice = ctx.QueryArray("sPPrice")
	sid = ctx.QueryArray("sid")
	esId = ctx.QueryArray("esID")

	// 整合sid和quantity
	build := strings.Builder{}
	for i := 0; i < len(sid); i++ {
		build.WriteString(sid[i] + "*" + quantity[i] + " ")
	}
	sids := build.String()

	// 开启事务, 注意后续使用tx来执行SQL
	tx := db.DB.Begin()

	// 更新产品的库存
	for i := 0; i < len(epId); i++ {
		id, err1 := strconv.Atoi(epId[i])
		quan, err2 := strconv.Atoi(quantity[i])
		if err1 != nil || err2 != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
		_, err := dao.UpdateStock(id, quan, tx)
		if err != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
	}

	// 转换spPrice为订单每个商品的总价
	pprice := make([]int, 0, len(epId))
	for i := 0; i < len(epId); i++ {
		pp, err1 := strconv.Atoi(spPrice[i])
		quan, err2 := strconv.Atoi(quantity[i])
		if err1 != nil || err2 != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
		pprice = append(pprice, quan*pp)
	}

	// 生成订单
	validPrice, err := strconv.Atoi(price)
	if err != nil {
		tx.Rollback()
		ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
		return
	}
	orderId, err := dao.CreateOrder(user.Id, user.UserName, user.Address, validPrice, tx)
	if err != nil {
		tx.Rollback()
		ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
		return
	}
	logrus.Infof("Create Order with ID %d\n", orderId)

	// 生成订单详情
	for i := 0; i < len(epId); i++ {
		id, err1 := strconv.Atoi(epId[i])
		quan, err2 := strconv.Atoi(quantity[i])
		if err1 != nil || err2 != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
		_, err = dao.GenerateOrderDetail(orderId, id, quan, pprice[i], tx)
		if err != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
	}

	// 更新购物车条目的状态
	for i := 0; i < len(esId); i++ {
		id, err := strconv.Atoi(esId[i])
		if err != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
		_, err = dao.SettleItem(id, tx)
		if err != nil {
			tx.Rollback()
			ctx.HTML(http.StatusOK, "shop_settle_err.html", gin.H{})
			return
		}
	}

	// 提交事务，并进行数据返回
	tx.Commit()
	sess.Set("WIDout_trade_no", epId)
	sess.Set("WIDtotal_amount", price)
	sess.Set("WIDbody", sids)
	sess.Save()
	attributes := gin.H{}
	attributes["WIDout_trade_no"] = epId
	attributes["WIDtotal_amount"] = price
	attributes["WIDbody"] = sids
	attributes["name"] = user
	ctx.HTML(http.StatusOK, "payOwn.tmpl", attributes)
}
