package redis

import (
	"bulebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// 计算起始，结束位置，并查询数据（封装）
func getIDsFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1

	// 查询 按照分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostIdsInOrder 从redis获取 id
func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	// 根据用户请求参数获取order中的参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 确定索引的起始位置
	// 查询 按照分数从大到小的顺序查询指定数量的元素
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 获取帖子的投票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	// 第一种方法
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	v := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	// 第二种方法,使用pipeline一次发送多条命令
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmders := range cmders {
		v := cmders.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIdsInOrder根据社区查询
func GetCommunityPostIdsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用zinterstore把分区的帖子set与按分数排序的zset生成一个新的zset
	// 针对新的zcet按照之前的逻辑取出数据
	// 利用缓存key减少zinterstore执行的次数（因为比较费时）

	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID))) // 社区的key
	key := orderKey + strconv.Itoa(int(p.CommunityID))                        // 缓存的key
	if rdb.Exists(orderKey).Val() < 1 {
		// 不存在key,需要计算
		piepline := rdb.Pipeline()
		piepline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)                   // ZInterStore 计算
		piepline.Expire(key, 60*time.Second) // 设置超时时间
		_, err := piepline.Exec()
		if err != nil {
			return nil, err
		}
	}

	return getIDsFromKey(key, p.Page, p.Size)
}
