package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"
	"goplay/model"
)

func Index(c *gin.Context) {
	/*session := sessions.Default(c)
	session.Set("username", "itcp")
	session.Save()
	//session.SetSession(c, "username", "itcp")
	username := session.Get("username")*/


	c.AsciiJSON(http.StatusOK, "ok")
}


func Addt(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("username")
	c.AsciiJSON(http.StatusOK, username)
}

func AddUser(c *gin.Context) {

	model.AddUser("itcp", "itcp1234", "my@itcp.name")
	c.AsciiJSON(http.StatusOK, "ok")
}