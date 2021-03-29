package controller

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"tmsshopping/dao"
)

type LoginForm struct {
	Username string `form:"userName"`
	Password string `form:"passWord"`
	// validate code
	//VeryCode string
}

func Login(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	// todo: 验证码验证
	var form LoginForm
	if ctx.ShouldBind(&form) == nil {
		//logrus.Infof("username:%s, password:%s", form.Username, form.Password)
		user, err := dao.SelectUserByNM(form.Username, form.Password)
		if err == nil {
			//session := sessions.Default(ctx)
			session := ginsession.FromContext(ctx)
			session.Set("name", user)
			session.Save()
			if user.Status == 2 {
				ctx.Redirect(http.StatusFound, "manage/indexPage")
			} else {
				ctx.Redirect(http.StatusFound, "indexSelect")
			}
		} else {
			ctx.HTML(http.StatusOK, "login_user_err.html", gin.H{})
		}
	} else {
		ctx.HTML(http.StatusOK, "login_err.html", gin.H{})
	}
}
