package serializer

import "singo/model"

// Box 快递柜信息序列化器
type Box struct {
	ID     int64  `json:"id"`
	IsUsed int32  `json:"is_used"`
	Key    string `json:"key"`
}

// BuildNewBoxSuccessRes 打包获取成功响应
func BuildNewBoxSuccessRes(box model.Box) Response {
	return Response{
		Data:  Box{
			ID:     box.Id,
			Key:    box.Key,
		},
		Msg:   "获取新的外卖柜成功",
	}
}

// BuildNoDataRes 无空闲外卖柜相应
func BuildNoDataRes(msg string) Response {
	return Response{
		Code: -1,
		Msg:   msg,
	}
}
