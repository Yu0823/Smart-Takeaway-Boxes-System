package service

import (
	"gorm.io/gorm"
	"singo/control"
	"singo/model"
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

// GetNewBoxService 获取一个新的外卖柜
type GetNewBoxService struct {
	BoxKey string `form:"box_key" json:"box_key" binding:"required"`
}

// GetNewBox 获取新外卖柜
func (service *GetNewBoxService) GetNewBox(c *gin.Context) serializer.Response {
	var box model.Box
	err := model.DB.Where("is_used = 0").First(&box).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return serializer.BuildNoDataRes("目前没有空闲的外卖柜！")
		} else{
			return serializer.Err(serializer.CodeDBError, "数据库操作失败", err)
		}
	}

	err = control.OpenBoxById(box.Id, false)
	if err != nil{
		return serializer.Err(serializer.CodeHardwareError, "硬件操作失败", err)
	}
	//获取后进行处理
	box.IsUsed = 1
	box.Key = service.BoxKey

	err = model.DB.Save(box).Error
	if err != nil{
		return serializer.Err(serializer.CodeDBError, "数据库操作失败", err)
	}

	return serializer.BuildNewBoxSuccessRes(box)
}
