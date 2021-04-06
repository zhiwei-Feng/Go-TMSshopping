package manage

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"time"
	"tmsshopping/db"
	"tmsshopping/domain"
)

func UserAddPage(ctx *gin.Context) {
	attributes := gin.H{}
	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	user, ok := loginUser.(domain.User)
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	attributes["name"] = user
	ctx.HTML(http.StatusOK, "user-add.tmpl", attributes)
}

// 前端调用的是GET Method
func UserAdd(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")

	// get params
	var (
		username string
		name     string
		pwd      string
		sex      string
		year     string
		email    string
		mobile   string
		address  string
	)

	username = ctx.Query("userName")
	name = ctx.Query("name")
	pwd = ctx.Query("passWord")
	sex = ctx.Query("sex")
	year = ctx.Query("birthday")
	email = ctx.Query("email")
	mobile = ctx.Query("mobile")
	address = ctx.Query("address")

	// parse birthday
	birthDay, err := time.Parse("2006-01-02", year)
	if err != nil {
		birthDay = time.Now()
	}

	newUser := domain.User{
		Id:       username,
		UserName: name,
		Password: pwd,
		Sex:      sex,
		Birthday: birthDay,
		Email:    email,
		Mobile:   mobile,
		Address:  address,
		Status:   1,
	}

	result := db.DB.Create(&newUser)
	if result.Error != nil {
		ctx.HTML(http.StatusOK, "user_add_err.html", gin.H{})
	}

	ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
}
