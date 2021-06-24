package logic

import (
	"bulebell/dao/mysql"
	"bulebell/dao/redis"
	"bulebell/models"
	"bulebell/pkg/snowflake"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	// 1.生成post id
	p.ID = snowflake.GetID()

	// 2.入库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID,p.CommunityID)
	if err != nil {
		return err
	}
	return
}

// GetPostDetail 根据id获取帖子详情
func GetPostDetail(id int64) (data *models.PostDetail, err error) {
	data = new(models.PostDetail)
	// 查询并拼接我们想要的数据
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetail() filed", zap.Error(err))
		return
	}
	// 根据作者ID查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() filed", zap.Error(err))
		return
	}

	// 根据社区ID查询社区姓名
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail() filed", zap.Error(err))
		return
	}

	data = &models.PostDetail{
		AuthorName: user.UserName,
		Post:       post,
		Community:  community,
	}

	// 获取数据
	return data, err
}

// GetPostList 获取帖子数据
func GetPostList(page, size int64) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	data = make([]*models.PostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() filed", zap.Error(err))
			continue
		}

		// 根据社区ID查询社区姓名
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail() filed", zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName: user.UserName,
			Post:       post,
			Community:  community,
		}
		data = append(data, postDetail)
	}

	return
}

// GetPostList 获取帖子数据
func GetPostList2(p *models.ParamPostList) (data []*models.PostDetail, err error) {
	// 去redis中查询id列表
	ids, err := redis.GetPostIdsInOrder(p)
	if err != nil {
		return
	}
	//fmt.Println("ids==",ids)
	if len(ids) == 0 {
		zap.L().Warn("len(ids) == 0")
		return
	}

	// 根据id去MySQL数据库查询帖子详情
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	// 提前查询好每篇帖子的投票数据
	votedata, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	for idx, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() filed", zap.Error(err))
			continue
		}

		// 根据社区ID查询社区姓名
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail() filed", zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName: user.UserName,
			Post:       post,
			Community:  community,
			VoteNum:    votedata[idx],
		}
		data = append(data, postDetail)
	}
	return
}

// GetCommunityPostList根据社区寻找帖子
func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.PostDetail, err error) {
	// 去redis中查询id列表
	ids, err := redis.GetCommunityPostIdsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("len(ids) == 0")
		return
	}
	// 根据id去MySQL数据库查询帖子详情
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// 提前查询好每篇帖子的投票数据
	votedata, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	for idx, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() filed", zap.Error(err))
			continue
		}
		// 根据社区ID查询社区姓名
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail() filed", zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName: user.UserName,
			Post:       post,
			Community:  community,
			VoteNum:    votedata[idx],
		}
		data = append(data, postDetail)
	}
	return
}
