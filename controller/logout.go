package controller

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Logout(ctx *gin.Context) {
	logrus.Info("Logout!")
	sess := ginsession.FromContext(ctx)
	sess.Delete("name")
	ctx.Redirect(http.StatusFound, "indexSelect")
}
