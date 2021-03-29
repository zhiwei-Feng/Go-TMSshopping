package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"os"
	"tmsshopping/controller"
	"tmsshopping/db"
)

func Logger() *logrus.Logger {
	//实例化
	logger := logrus.New()
	logger.Out = os.Stdout

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func add(x, y int) int {
	return x + y
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

	// +--------------+ site config
	Logger().Infoln("Welcome to TMS shopping application")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/tms?charset=utf8&parseTime=True&loc=Local", mysqlu, mysqlp, mysqlAddr)
	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db conn fail.")
	} else {
		Logger().Infoln("database connect success.")
		db.DB = dbconn
	}

	// +--------------+ gin
	router := gin.Default()
	router.Use(gin.Recovery())
	// gin-contrib/sessions就是不好使, 所以只能用第三方库go-session/gin-session
	router.Use(ginsession.New())
	// 定义一些模板函数
	router.SetFuncMap(template.FuncMap{"add": add})
	// 加载模板及html文件，注意这种写法下，templates根目录的模板无法加载
	router.LoadHTMLGlob("templates/**/*")
	// 加载静态文件
	router.Static("/css", "./static/css")
	router.Static("/images", "./static/images")
	router.Static("/scripts", "./static/scripts")
	// +--------------+ 静态页面渲染
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/loginPage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})
	// +--------------+ http请求
	router.GET("/indexSelect", controller.IndexSelect)
	router.GET("/selectProductList", controller.SelectProductList)
	router.GET("/selectProductView", controller.SelectProductView)
	router.POST("/login", controller.Login)
	router.GET("/zx", controller.Logout)
	router.GET("/SelallServlet", controller.MessageBoard)
	_ = endless.ListenAndServe(":8888", router)
}
