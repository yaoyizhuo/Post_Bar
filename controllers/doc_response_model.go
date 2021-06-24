package controllers

import "bulebell/models"

// 存放接口文档用到的model
// 因为返回数据的格式是一致的，但是具体的data类型不一致
type _ResponsePostList struct {
	Code    ResCode        `json:"code"`
	Message string         `json:"message"`
	Data    []*models.User `json:"data"`
}
