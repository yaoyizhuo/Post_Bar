package redis

// 存放redis的key值
// redis key 注意命名方式，方便差分和查询
const (
	KeyPrefix        = "bulebell:"
	KeyPostTimeZSet  = "post:time"   // 帖子以及发帖时间
	KeyPostScoreZSet = "post:score"  // 帖子以及投票分书
	KeyPostVotedZSetPF     = "post:voted:" // 记录用户以及投票类型   需要的参数是post ID
	KeyCommunitySetPF = "community:"	//保存每个帖子的分区
)

func getRedisKey(key string)string  {
	return KeyPrefix+key
}