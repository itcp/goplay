package cbase

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"reflect"
	//"goplay/public/common"
)

func UploadImg(c *gin.Context) {
	Response := map[string]interface{}{
		"status": 0,
		"info": "",
		"path": "",
	}
	file, _ := c.FormFile("file")
	fmt.Println(reflect.TypeOf(file))
	//fmt.Println(reflect.TypeOf(file))
	theImgPath := "./static/imwf45.jpg"
	c.SaveUploadedFile(file, theImgPath)
	Response["path"] = theImgPath
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