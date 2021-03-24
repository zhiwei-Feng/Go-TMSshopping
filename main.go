package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func main() {
	mysqlu := flag.String("mysqlu", "root", "mysql user")
	mysqlp := flag.String("mysqlp", "root", "mysql password")
	mysqlAddr := flag.String("server", "127.0.0.1", "mysql server address")

	flag.Parse()

	// +--------------+ site config
	Logger().Infoln("Welcome to TMS shopping application")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/tms?charset=utf8&parseTime=True&loc=Local", *mysqlu, *mysqlp, *mysqlAddr)
	//dsn := "root:feng1995@tcp(10.176.64.25:3306)/tms?charset=utf8&parseTime=True&loc=Local"
	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db conn fail.")
	} else {
		Logger().Infoln("database connect success.")
		db.DB = dbconn
	}

	store := cookie.NewStore([]byte("secret"))

	// +--------------+ gin
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("mySession", store))
	router.LoadHTMLGlob("templates/*")
	router.Static("/css", "./static/css")
	router.Static("/images", "./static/images")
	router.Static("/scripts", "./static/scripts")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/indexSelect", controller.IndexSelect)
	router.GET("/product-view", controller.SelectProductView)
	_ = endless.ListenAndServe(":8080", router)
}
