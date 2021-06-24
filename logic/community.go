package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
)

// CommunityHandler 查询分类信息
func GetCommunityList() ([]*models.Community, error) {
	// 从数据库获取数据
	return mysql.GetCommunityList()
}

// CommunityByIDHandler 根据ID查询详细信息
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	// 从数据库获取数据
	return mysql.GetCommunityDetail(id)
}
