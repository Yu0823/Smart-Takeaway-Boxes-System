package model

import (
	"gorm.io/gorm"
)

// Box 快递柜模型
type Box struct {
	gorm.Model
	Id     int64
	IsUsed int32
	Key    string
}

// GetBoxById 用id获取对应box
func GetBoxById(ID interface{}) (Box, error) {
	var box Box
	result := DB.First(&box, ID)
	return box, result.Error
}
