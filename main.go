package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"os"
	"tmsshopping/controller"
	"tmsshopping/controller/manage"
	"tmsshopping/controller/shop"
	"tmsshopping/db"
	"tmsshopping/log"
)

func add(x, y int) int {
	return x + y
}

func reduce(x, y int) int {
	return x - y
}

func main() {
	mysqlu, exist := os.LookupEnv("MYSQLU")
	if !exist {
		mysqlu = "root"
	}
	mysqlp, exist := os.LookupEnv("MYSQLP")
	if !exist {
		mysqlp = "root"
	}
	mysqlAddr, exist := os.LookupEnv("MYSQL_ADDR")
	if !exist {
		mysqlAddr = "127.0.0.1"
	}
	database, exist := os.LookupEnv("DB")
	if !exist {
		database = "fzw"
	}

	// +--------------+ site config
	// db
	log.Log.Infoln("Welcome to TMS shopping application")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", mysqlu, mysqlp, mysqlAddr, database)
	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db conn fail.")
	} else {
		log.Log.Debugln("database connect success.")
		db.DB = dbconn
	}

	// +--------------+ gin
	router := gin.Default()
	router.Use(gin.Recovery())
	// gin-contrib/sessions就是不好使, 所以只能用第三方库go-session/gin-session
	router.Use(ginsession.New())
	// 定义一些模板函数
	router.SetFuncMap(template.FuncMap{"add": add, "reduce": reduce})
	// 加载模板及html文件，注意这种写法下，templates根目录的模板无法加载
	router.LoadHTMLGlob("templates/**/*")
	// 加载静态文件
	router.Static("/css", "./static/css")
	router.Static("/images", "./static/images")
	router.Static("/scripts", "./static/scripts")
	// +--------------+ 静态页面渲染
	router.GET("/test", func(context *gin.Context) {
		log.Log.WithField("test", "test").Info("测试文本")
		log.Log.WithField("warn", "warn").Warn("警告文本")
		log.Log.WithField("debug", "debug").Debug("警告文本")
		log.Log.WithField("error", "error").Error("错误文本")
	})
	router.GET("/loginPage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})
	router.GET("/registerPage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.tmpl", gin.H{})
	})
	router.GET("/registerResultPage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "reg-result.tmpl", gin.H{})
	})
	// +--------------+ http请求
	router.GET("/", controller.IndexSelect)
	router.GET("/index", controller.IndexSelect)
	router.GET("/indexSelect", controller.IndexSelect)             //首页
	router.GET("/selectProductList", controller.SelectProductList) // 商品列表
	router.GET("/selectProductView", controller.SelectProductView) // 商品详情页
	router.GET("/zx", controller.Logout)                           // 用户登出
	router.GET("/SelallServlet", controller.MessageBoard)          // 留言
	router.GET("/usernamecheck", controller.UsernameCheck)         // 注册用户时验证用户名
	router.GET("/ShopSelect", shopController.ShopSelect)           // 购物车
	router.GET("/shopAdd", shopController.ShopBuy)                 // 购买商品按钮
	router.GET("/shopAdd2", shopController.ShopAdd)                // 放入购物车
	router.GET("/selectdd", controller.SelectOrder)                // 个人订单
	router.GET("/UpdateServlet", shopController.ShopUpdate)        // 购物车购买数量更新
	router.GET("/gmServlet", shopController.ShopCartSettle)        // 购物车结算
	router.POST("/login", controller.Login)                        // 登录
	router.POST("/GueServlet", controller.PostComment)             // 提交留言
	router.POST("/register", controller.Register)                  // 注册

	// +--------------+ manage part
	m := router.Group("/manage")
	m.GET("/index", manage.Index)
	m.GET("/userAddPage", manage.UserAddPage)
	m.GET("/useradd", manage.UserAdd)
	m.GET("/user", manage.UserManagePage)
	m.GET("/userdel", manage.UserDelete)
	m.GET("/userUpdatePage", manage.UserUpdatePage)
	m.GET("/userupdate", manage.UserUpdate)
	m.GET("/productClass", manage.ProductClassManagePage)
	m.GET("/productClassAdd", manage.ProductClassAddPage)
	m.GET("/doProductClassAdd", manage.ProductClassAdd)
	m.GET("/productClassDel", manage.ProductClassDel)
	m.GET("/toProductClassUpdate", manage.ProductClassUpdatePage)
	m.GET("/doProductClassUpdate", manage.ProductClassUpdate)
	m.GET("/productSelect", manage.ProductManagePage)
	m.GET("/productAdd", manage.ProductAddPage)
	m.GET("/ordersel", manage.OrderManagePage)
	m.GET("/delorder", manage.OrderDelete)
	m.GET("/touporder", manage.OrderUpdatePage)

	_ = endless.ListenAndServe(":8888", router)
}
