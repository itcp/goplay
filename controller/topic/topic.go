package topic

import(
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.HTML(http.StatusOK, "m_addt.html", gin.H{})
}

func AddTopicApi(c *gin.Context) {
	c.JSONP(http.StatusOK, "提交成功！")
}
