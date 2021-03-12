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