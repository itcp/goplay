package user

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "m_register.html", gin.H{})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "m_login.html", gin.H{})
}

func Logout(c *gin.Context) {
	c.JSONP(http.StatusOK, "退出登录")
}