package topic

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"goplay/model"
	"github.com/gin-contrib/sessions"
	"strconv"
	"strings"
	"fmt"
)

func Index(c *gin.Context) {
	topicList, _ := model.GetPageTopic(0, 0)
	typeList, _ := model.GetTopicTypeList()
	c.HTML(http.StatusOK, "m_index.html", gin.H{
		"typeList": typeList,
		"topicList": topicList,
	})
}

// 列表
func List(c *gin.Context) {
	tid,_ := strconv.Atoi(c.DefaultQuery("tid", "0"))
	page,_ := strconv.Atoi(c.DefaultQuery("page", "0"))
	topicList, _ := model.GetPageTopic(tid, page)
	typeList, _ := model.GetTopicTypeList()
	theTypeName := ""
	for _, tmap := range typeList{
		if tmap.ID == tid{
			theTypeName = tmap.Name
		}
	}
	c.HTML(http.StatusOK, "m_tlist.html", gin.H{
		"typeList": typeList,
		"theTypeName": theTypeName,
		"topicList": topicList,
		"page": page,
		"tid": tid,
		"lastPage": page - 1,
		"nextPage": page + 1,
	})
}
// 详情页
func Details(c *gin.Context) {
	tid := c.Param("tid")
	fmt.Println(tid)
	theTopic, _ := model.GetTopicByID(tid)
	topicImgList, _ := model.GetTopicImgByTid(tid)
	typeList, _ := model.GetTopicTypeList()
	c.HTML(http.StatusOK, "m_tcont.html", gin.H{
		"typeList": typeList,
		"theTopic": theTopic,
		"timgList": topicImgList,
})
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
	imglist := c.PostForm("imglist")
	
	aImglist := strings.Split(imglist, ",")

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
        Mimg: aImglist[0],
        Status: 1,
	}
	theTid, err := model.AddTopic(theTopic)

	if err != nil {
		Response["status"] = 2
	}else{
		for _, filepath := range aImglist {
			timg := &model.Timg{
				Tmid: theTid,
				Path: filepath,
			}
			model.AddTopicImg(timg)
		}
		Response["status"] = 1
	}
	
	c.JSONP(http.StatusOK, Response)
}

func PassTopic(c *gin.Context) {
	topicList, _ := model.GetPassTopic()
	c.HTML(http.StatusOK, "m_ptlist.html", gin.H{
		"topicList": topicList,
	})
}

func PassTopicApi(c *gin.Context) {
	Response := map[string]interface{}{
		"status": 0,
		"info": "",
	}
	ptid := c.PostForm("ptid")
	pstatus,_ := strconv.Atoi(c.PostForm("pstatus"))
	err := model.PassTopic(ptid, pstatus)
	if err != nil{
		Response["status"] = 2
		return
	}
	Response["status"] = 1
	c.JSONP(http.StatusOK, Response)
}