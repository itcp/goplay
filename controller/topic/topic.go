package topic

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"goplay/model"
	"github.com/gin-contrib/sessions"
	"strconv"
	"fmt"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "m_index.html", gin.H{})
}

// 列表
func List(c *gin.Context) {
	c.HTML(http.StatusOK, "m_tlist.html", gin.H{})
}
// 详情页
func Details(c *gin.Context) {
	c.HTML(http.StatusOK, "m_tcont.html", gin.H{})
}
// 信息发布
func AddTopic(c *gin.Context) {
	TopicTypeList, _ := model.GetTopicTypeList()
	c.HTML(http.StatusOK, "m_addt.html", gin.H{"TopicTypeList": TopicTypeList,})
}

func AddTopicApi(c *gin.Context) {
	Response := map[string]interface{}{
		"status": 0,
		"info": "",
	}
	title := c.PostForm("title")
	ttid,_ := strconv.Atoi(c.PostForm("type"))
	content := c.PostForm("content")
	mimg := c.PostForm("mimg")
	
    // 先到数据库查询用户名是否已用,已用的状态码是 status=3
	titleIn := model.ExistTopicByTitle(title)
	if titleIn {
		Response["status"] = 3
	}
	session := sessions.Default(c)
	theTopic := &model.Topic{
		Ttid:    ttid,
		Uid:     session.Get("uid").(string),
		Title: title,
        Content: content,
        Mimg: mimg,
        Status: 1,
	}
	fmt.Println(theTopic)
	err := model.AddTopic(theTopic)
	if err != nil {
		Response["status"] = 2
	}else{
		Response["status"] = 1
	}
	
	c.JSONP(http.StatusOK, Response)
}
