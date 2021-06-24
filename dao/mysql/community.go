package mysql

import (
	"bulebell/models"
	"database/sql"
	"go.uber.org/zap"
)

// CommunityHandler 查询分类信息
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`

	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = ErrInvalidID
		}
	}
	//fmt.Println(communityList)
	return
}

// CommunityByIDHandler 根据ID查询详细信息
func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail) // 类型初始化
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id=?`

	if err := db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidID
		}
	}
	return communityDetail, err
}

// CommunityByIDHandler 根据ID查询名字
func GetCommunityByID(id int64) (community *models.Community, err error) {
	community = new(models.Community) // 类型初始化
	sqlStr := `select community_id,community_name from community where community_id=?`

	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidID
		}
	}

	return
}
