package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

/*	投票的几种情况
direction=1：
	1.之前没有投过票，现在偷赞成票		-->更新分数和投票记录		差值的绝对值为 1	+432
	2.之前反对票，现在改偷赞成票		-->更新分数和投票记录		差值的绝对值为 2	+432*2
direction=0：
	1.之前为赞成票，现在取消			-->更新分数和投票记录		差值的绝对值为 1	-432
	2.之前为反对票，现在取消			-->更新分数和投票记录		差值的绝对值为 1	+432
direction=-1：
	1.之前没有投过票，现在偷反对票		-->更新分数和投票记录		差值的绝对值为 1	-432
	2.之前赞成票，现在改偷反对票		-->更新分数和投票记录		差值的绝对值为 2	-432*2

投票的限制：
	每个帖子在发帖一周之内允许用户投票，超过一星期就不行了
	1.到期之后将redis中保存的赞成和反对票存储到mysql
	2.到期之后删除 redis中的保存记录
*/
const (
	oneWeek      = 7 * 24 * 3600
	scorePerVote = 432 //每一票的分数为432分。
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepeated = errors.New("不允许重复投票")
)

// CreatePost 创建帖子时间
func CreatePost(postID,communityID int64) (err error) {
	//事务，要么同时成功，要么都不成功
	pipeline := rdb.Pipeline()

	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 补充：把帖子ID加入到社区的set
	cKey := getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey,postID)

	_, err = pipeline.Exec()

	return err
}

// VoteForPost 投票逻辑
func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票限制
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeek {
		return ErrVoteTimeExpire
	}

	// 2.更新帖子分数
	// 查询当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	// 如果当前投票和之前用户对他的投票一样，就提示不允许重复投票
	if value == ov{
		return ErrVoteRepeated
	}
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value) // 计算投票的差值

	pipeline := rdb.Pipeline() // 第二部和第三步要放入一个事务
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)

	// 3.记录用户对改帖子投的票数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)

	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,  // 赞成票还是反对票
			Member: userID, // 哪个用户
		})
	}

	_, err := pipeline.Exec()
	return err
}
