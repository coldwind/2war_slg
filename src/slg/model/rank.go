package model

import (
	"slg/common"

	//	"github.com/garyburd/redigo/redis"
)

type Rank struct {
}

func (this *Rank) SetRank(uid int, score int) {

	redisHandle := common.RedisHandle.GetRedisHandle(0)
	defer redisHandle.Close()

	key := "rank:0"
	_, _ = redisHandle.Do("zadd", key, score, uid)
}

func (this *Rank) GetRankList() {
	//redisHandle := common.RedisHandle.GetRedisHandle(0)
	//res, err := redisHandle.Do("")
}
