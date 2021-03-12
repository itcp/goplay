package user

import(
	"github.com/gin-gonic/gin"
	"net/http"
	//"fmt"
	"goplay/model"
	"github.com/gin-contrib/sessions"
	"goplay/public/common"
)

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "m_register.html", gin.H{})
}

func RegisterApi(c *gin.Context) {
	Response := map[string]interface{}{
		"status": "0",
		"info": "",
	}
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
    // 先到数据库查询用户名是否已用,已用的状态码是 status=3
	userIn := model.ExistUserByName(username)
	
	if userIn {
		// 存到数据库成功, 返回状态码是 status =1
		_, err := model.AddUser(username, password, email)
		if err != nil {
			Response["status"] = 2
			return
		}
		Response["status"] = 1
	}else{
		Response["status"] = 3
	}

	c.JSONP(http.StatusOK, Response)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "m_login.html", gin.H{})
}

func LoginApi(c *gin.Context) {
	Response := map[string]interface{}{
		"status": "0",
		"info": "",
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
    // 先到数据库查询用户名是否已用,已用的状态码是 status=3
	theuser, err := model.GetUserByName(username)
	if err != nil{
		Response["status"] = 3
		return
	}
	if common.VerifyString(password,theuser.Password ){
		session := sessions.Default(c)
		session.Set("username", username)
		session.Save()
		Response["status"] = 1
	}

	c.JSONP(http.StatusOK, Response)
}

func Logout(c *gin.Context) {
	c.JSONP(http.StatusOK, "退出登录")
}