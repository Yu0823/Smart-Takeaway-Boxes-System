package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// GetNewBox 获取新外卖柜接口
func GetNewBox (c *gin.Context) {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("ctx.Request.body: %v\n", string(data))
	var service service.GetNewBoxService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNewBox(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//OpenBoxById 根据id打开外卖柜
func OpenBoxById (c *gin.Context) {
	var service service.OpenBoxByIdService
	if err := c.ShouldBind(&service); err == nil {
		res := service.OpenBoxById(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}