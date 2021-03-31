package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tmsshopping/dao"
)

// GET method
func UsernameCheck(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	name := ctx.Query("name")
	if strings.TrimSpace(name) == "" {
		ctx.String(http.StatusOK, "false")
		return
	}
	_, err := dao.SelectUserByName(name)
	if err == nil {
		ctx.String(http.StatusOK, "false")
	} else {
		ctx.String(http.StatusOK, "true")
	}
}
