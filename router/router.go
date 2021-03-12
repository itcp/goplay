package router

import (
	"goplay/controller"
	"goplay/controller/user"
	"goplay/controller/topic"
	"goplay/middleware/auto"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	store := cookie.NewStore([]byte("secret11111"))
	r.Use(sessions.Sessions("mysession", store))
	// 引入session

	// 推荐使用绝对路径 相当于简历了软连接--快捷方式
	//r.StaticFS("/static", http.Dir("./static"))
	//r.StaticFS("/upload", http.Dir("./upload"))

	// 未登录可访问的url
	r.LoadHTMLGlob("templates/*/**")
	pubweb := r.Group("")
	{
		// 首页
		pubweb.GET("/", controller.Index)
		pubweb.GET("/adduser", controller.AddUser)

		pubweb.GET("/register", user.Register)
		pubweb.POST("/register", user.RegisterApi)
		pubweb.GET("/login", user.Login)
		pubweb.POST("/login", user.LoginApi)

		pubweb.GET("/tlist", topic.List)
		pubweb.GET("/tdetails", topic.Details)
		pubweb.GET("/add", topic.AddTopic)

	}
	pubweb.Use(auto.LoginCheck())
	{
		//pubweb.GET("/add", topic.AddTopic)	
		pubweb.POST("/add", topic.AddTopicApi)
	}

	userweb := r.Group("/user")
	userweb.Use(auto.LoginCheck())
	{
		// 信息发布
		userweb.GET("/addt", controller.Addt)	
	}

	return r
}
