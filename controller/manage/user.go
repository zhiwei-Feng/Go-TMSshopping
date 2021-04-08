package manage

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"net/http"
	"strconv"
	"time"
	"tmsshopping/dao"
	"tmsshopping/db"
	"tmsshopping/domain"
	"tmsshopping/log"
)

// 新增用户页面
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

// 用户管理页面
func UserManagePage(ctx *gin.Context) {
	var (
		cpage      = 1
		count      = 5
		cp         = ctx.Query("cp")
		attributes = gin.H{}
	)

	if v, err := strconv.Atoi(cp); err == nil {
		cpage = v
	}

	tpage, _ := dao.TotalPageForUser(int64(count))
	list, _ := dao.UserPagination(cpage, count)

	selectList := make([]int, 0, tpage)
	for i := 1; i <= tpage; i++ {
		selectList = append(selectList, i)
	}

	attributes["userlist"] = list
	attributes["cpage"] = cpage
	attributes["tpage"] = tpage
	attributes["selectList"] = selectList
	ctx.HTML(http.StatusOK, "user-manage.tmpl", attributes)
}

// 用户删除 GET Method
func UserDelete(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		id = ctx.Query("id")
	)

	delUser := domain.User{Id: id}
	result := db.DB.Delete(&delUser)
	if result.Error != nil {
		ctx.HTML(http.StatusOK, "user_del_err.html", gin.H{})
	}

	ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
}

// 用户修改页面
func UserUpdatePage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		id         = ctx.Query("id")
		attributes = gin.H{}
	)
	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	user, ok := loginUser.(domain.User)
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
	}
	updateUser, _ := dao.SelectUserByName(id)

	attributes["name"] = user
	attributes["user"] = updateUser

	ctx.HTML(http.StatusOK, "user-modify.tmpl", attributes)
}

// 用户修改功能
func UserUpdate(ctx *gin.Context) {
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
	birthDay, err := time.Parse("2006-01-02 15:04:05 +0800 CST", year)
	if err != nil {
		birthDay, err = time.Parse("2006-1-2", year)
		if err != nil {
			log.Log.WithField("err", err).Warn("生日格式转换错误")
			ctx.HTML(http.StatusInternalServerError, "user_mod_err.html", gin.H{})
			return
		}
	}

	updateUser := domain.User{
		Id:       username,
		UserName: name,
		Password: pwd,
		Sex:      sex,
		Birthday: birthDay,
		Email:    email,
		Mobile:   mobile,
		Address:  address,
	}

	result := db.DB.Model(&updateUser).Omit("EU_STATUS", "EU_USER_NAME").Updates(updateUser)
	if result.Error == nil {
		ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
		return
	}

	log.Log.WithField("err", result.Error.Error()).Warn("用户更新失败")
	ctx.HTML(http.StatusInternalServerError, "user_mod_err.html", gin.H{})
}
