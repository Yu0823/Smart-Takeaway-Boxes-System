package service

import (
	"fmt"
	"gorm.io/gorm"
	"singo/model"
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

// OpenBoxByIdService 获取一个新的外卖柜
type OpenBoxByIdService struct {
	Id int64 `form:"box_id" json:"box_id" binding:"required"`
	BoxKey string `form:"box_key" json:"box_key" binding:"required"`
}

// OpenBoxById 通过id和密钥获取打开外卖柜
func (service *OpenBoxByIdService) OpenBoxById(c *gin.Context) serializer.Response {
	var box model.Box
	err := model.DB.Where("id = ?", service.Id).First(&box).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return serializer.BuildNoDataRes("输入的外卖柜id有误！")
		} else{
			return serializer.Err(serializer.CodeDBError, "数据库操作失败", err)
		}
	}

	//如果外卖柜里没东西
	if box.IsUsed == 0 {
		return serializer.BuildNoDataRes("此外卖柜为空！")
	}
	//如果密钥不对
	if service.BoxKey != box.Key {
		return serializer.BuildNoDataRes("此id对应的密钥错误！")
	}

	//开门后进行处理
	box.IsUsed = 0
	box.Key = ""

	err = model.DB.Save(box).Error
	if err != nil{
		return serializer.Err(serializer.CodeDBError, "数据库操作失败", err)
	}

	return serializer.BuildNoDataRes(fmt.Sprintf("%d号外卖柜开门成功",box.Id))
}
