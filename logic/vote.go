package logic

import (
	"bulebell/dao/redis"
	"bulebell/models"
	"go.uber.org/zap"
	"strconv"
)

/*
1.用户投票的数据
2.http://www.ruanyifeng.com/blog/algorithm/
*/

// 投票功能:
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))

	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
