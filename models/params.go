// 与请求相关的参数模型
package models

const (
	OrderTime = "time"
	OrderScore = "score"

)


// 定义请求参数结构体
type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 登陆表单
type ParamsLoginUser struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              //帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票还是反对  1，-1。。。取消投票为 0
}

// ParamPostList获取帖子列表的各个参数（升级版）
type ParamPostList struct {
	Page  int64`json:"page" form:"page"`
	Size  int64`json:"size" form:"size"`
	Order string`json:"order" form:"order"`
}

// ParamCommunityPostList根据社区获取帖子列表的各个参数
type ParamCommunityPostList struct {
	ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}