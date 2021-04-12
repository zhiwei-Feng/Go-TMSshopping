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
)

func CommentManagePage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		page       = 1
		pageSize   = 3
		spage      = ctx.Query("page")
		list       []domain.Comment
		maxPage    int
		attributes = gin.H{}
	)

	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
		return
	}

	if v, err := strconv.Atoi(spage); err == nil {
		page = v
	}

	list, _ = dao.CommentPage(page, pageSize)
	maxPage, _ = dao.MaxCommentPageNum(int64(pageSize))

	// 页面所需数组
	plist := make([]int, 0, maxPage)
	for i := 1; i <= maxPage; i++ {
		plist = append(plist, i)
	}

	attributes["name"] = loginUser
	attributes["list"] = list
	attributes["max_page"] = maxPage
	attributes["page"] = page
	attributes["plist"] = plist

	ctx.HTML(http.StatusOK, "comment-manage.tmpl", attributes)
}

func CommentDelete(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		idStr = ctx.Query("id")
		id    int
		err   error
	)

	if id, err = strconv.Atoi(idStr); err != nil {
		ctx.HTML(http.StatusBadRequest, "comment_del_err.html", gin.H{})
		return
	}

	result := db.DB.Delete(&domain.Comment{}, id)
	if result.Error != nil {
		ctx.HTML(http.StatusInternalServerError, "comment_del_err.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
}

func CommentReplyPage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		idStr      = ctx.Query("id")
		id         int
		comment    domain.Comment
		attributes = gin.H{}
	)

	id, _ = strconv.Atoi(idStr)
	db.DB.First(&comment, id)

	sess := ginsession.FromContext(ctx)
	loginUser, ok := sess.Get("name")
	if !ok {
		ctx.HTML(http.StatusOK, "login_first.html", gin.H{})
		return
	}

	attributes["list"] = comment
	attributes["name"] = loginUser

	ctx.HTML(http.StatusOK, "comment-reply.tmpl", attributes)
}

func CommentReply(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		orderId      = ctx.PostForm("orderId")
		replyContent = ctx.PostForm("replyContent")
		id           int
		err          error
	)

	if id, err = strconv.Atoi(orderId); err != nil {
		ctx.HTML(http.StatusBadRequest, "comment_reply_err.html", gin.H{})
		return
	}

	upComment := domain.Comment{
		Id:        id,
		Reply:     replyContent,
		ReplyTime: time.Now(),
	}

	result := db.DB.Model(&upComment).Select("EC_REPLY", "EC_REPLY_TIME").Updates(upComment)
	if result.Error != nil {
		ctx.HTML(http.StatusInternalServerError, "comment_reply_err.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "manage-result.tmpl", gin.H{})
}
