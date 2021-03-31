package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"tmsshopping/db"
	"tmsshopping/domain"
)

func Register(ctx *gin.Context) {
	// todo: 增加验证码验证
	formUsername := ctx.PostForm("userName")
	formName := ctx.PostForm("name")
	formRePassword := ctx.PostForm("rePassWord")
	formSex := ctx.PostForm("sex")
	formBirthday := ctx.PostForm("birthday")
	formEmail := ctx.PostForm("email")
	formMobile := ctx.PostForm("mobile")
	formAddress := ctx.PostForm("address")

	if strings.TrimSpace(formUsername) == "" || strings.TrimSpace(formRePassword) == "" {
		ctx.HTML(http.StatusOK, "register_err.html", gin.H{})
	}

	// 时间转化
	validBirthday, err := time.Parse("2006-01-02", formBirthday)
	if err != nil {
		validBirthday = time.Now()
	}

	newUser := domain.User{
		Id:       formUsername,
		UserName: formName,
		Password: formRePassword,
		Sex:      formSex,
		Birthday: validBirthday,
		Email:    formEmail,
		Mobile:   formMobile,
		Address:  formAddress,
		Status:   1, // 暂时硬编码
	}

	result := db.DB.Create(&newUser)
	if result.Error != nil {
		ctx.HTML(http.StatusOK, "register_err.html", gin.H{})
	} else {
		ctx.Redirect(http.StatusFound, "reg-result.jsp")
	}
}
