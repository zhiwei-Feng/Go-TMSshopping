package main

import (
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
