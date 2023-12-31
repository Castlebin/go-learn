package runtime_panic

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/util"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				/*
					subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.AppName)

					body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
					body  = strings.ReplaceAll(body, "{RequestTime}", util.GetCurrentDate())
					body  = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method + "  " + c.Request.Host + c.Request.RequestURI)
					body  = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
					body  = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
					body  = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

					_ = util.SendMail(config.ErrorNotifyUser, subject, body)
				*/
				utilGin := util.Gin{Ctx: c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
