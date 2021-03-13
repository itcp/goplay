package cbase

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"os"
	"goplay/model"
)

func UploadImg(c *gin.Context) {
	Response := map[string]interface{}{
		"status": 0,
		"info": "",
		"path": "",
	}
	file, _ := c.FormFile("file")
	timeTemplate3 := "060102"
    now := time.Now()
    stday := now.Format(timeTemplate3) 
    yearm := stday[0:4]
    mday := stday[4:]
	theImgDir := "./static/images/" + yearm + "/" + mday
	err:=os.MkdirAll(theImgDir, os.ModePerm)
	if err != nil{
		Response["status"] = 2
		return
	}
	theImgPath := theImgDir + "/"+ model.GetID() + ".png"

	c.SaveUploadedFile(file, theImgPath)
	Response["path"] = theImgPath[1:]

	c.JSONP(http.StatusOK, Response)
}

// 分页数据结构体
type PaginationQ struct {
	Ok    bool        `json:"ok"`
	Page  uint        `form:"page" json:"page"`  // 请求的页码
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total uint        `json:"total"`
}
/*
func GetPagination(page, total int)(pageData PaginationQ){

}*/