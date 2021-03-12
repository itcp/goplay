package auto 

import (
	"github.com/gin-gonic/gin"
	"goplay/public/rcode"
	"net/http"
	"github.com/gin-contrib/sessions"
)
// 验证是否登录
func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		session := sessions.Default(c)
		username := session.Get("username")
		isLogin := false
		if username != nil{
			isLogin = true
		}
		

		if !isLogin {
			code = rcode.UNPASS
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  rcode.GetMessage(code),
			})
			c.Redirect(301, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
