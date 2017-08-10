package model

import (
	"slg/common"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type PveMap struct {
	PveMapOwnPrefix string // PVE拥有地图
	PveMapAtkPrefix string // 进攻信息
}

func NewPveMap() (pveMap *PveMap) {

	pveMap = &PveMap{}
	pveMap.PveMapOwnPrefix = "pvemap:own:"
	pveMap.PveMapAtkPrefix = "pvemap:atk:"

	return pveMap
}

func (this *PveMap) Add(uid int, mid int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)

	// 写入城市容器
	key := this.PveMapOwnPrefix + uidStr
	_, reserr := redisHandle.Do("sadd", key, mid)
	if reserr != nil {
		return false
	}

	return true
}

func (this *PveMap) GetAll(uid int) (mapId []int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)

	// 写入城市容器
	key := this.PveMapOwnPrefix + uidStr
	res, reserr := redisHandle.Do("smembers", key)
	mapId = make([]int, 0, 961)
	if reserr != nil {
		return mapId
	}
	mapId, _ = redis.Ints(res, reserr)

	return mapId
}

func (this *PveMap) Atk(uid int, corpsId int, atkInfo string) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	// 初始化基础信息
	key := this.PveMapAtkPrefix + strconv.Itoa(uid)
	res, reserr := redisHandle.Do("hset", key, corpsId, atkInfo)
	if reserr != nil {
		return false
	}

	_, err := redis.Int(res, reserr)

	if err != nil {
		return false
	}

	return true
}
