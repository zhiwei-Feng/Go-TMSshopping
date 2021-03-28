package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `form:"userName"`
	Password string `form:"passWord"`
	// validate code
	//VeryCode string
}

func Login(ctx *gin.Context) {
	// todo: 验证码验证
	var form LoginForm
	if ctx.ShouldBind(&form) == nil {
		//logrus.Infof("username:%s, password:%s", form.Username, form.Password)

	} else {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.HTML(http.StatusOK, "login_err.html", gin.H{})
	}
}
