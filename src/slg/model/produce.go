package model

import (
	"slg/common"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Produce struct {
	InfoPrefix   string
	FacSetPrefix string
}

func NewProduce() (produce *Produce) {

	produce = &Produce{}
	produce.InfoPrefix = "fac:info:"
	produce.FacSetPrefix = "fac:"

	return
}

func (this *Produce) Add(uid int, cid int, id int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	idStr := strconv.Itoa(id)
	cidStr := strconv.Itoa(cid)

	key := this.FacSetPrefix + cidStr + ":" + uidStr
	res, reserr := redisHandle.Do("hset", key, id, id)
	_, err := redis.Int(res, reserr)
	if err != nil {
		return false
	}

	now := time.Now().Unix()
	timeStr := strconv.Itoa(int(now))

	infoKey := this.InfoPrefix + uidStr + ":" + cidStr + ":" + idStr
	res, reserr = redisHandle.Do("hmset", infoKey, "id", idStr, "level", 1, "collectTime", timeStr)
	_, err = redis.Int(res, reserr)

	if err != nil {
		return false
	}

	return true
}

func (this *Produce) Set(uid int, cid int, id int, collectTime int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	idStr := strconv.Itoa(id)
	cidStr := strconv.Itoa(cid)

	key := this.InfoPrefix + uidStr + ":" + cidStr + ":" + idStr

	_, err := redisHandle.Do("hset", key, "collectTime", strconv.Itoa(collectTime))
	if err != nil {
		return false
	}

	return true
}

func (this *Produce) SetLevel(uid int, cid int, id int, level int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	idStr := strconv.Itoa(id)
	cidStr := strconv.Itoa(cid)

	key := this.InfoPrefix + uidStr + ":" + cidStr + ":" + idStr

	_, err := redisHandle.Do("hset", key, "level", level)
	if err != nil {

		return false
	}

	return true
}

func (this *Produce) Get(uid int, cid int, id int) (factoryInfo *common.Factory) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	idStr := strconv.Itoa(id)
	cidStr := strconv.Itoa(cid)

	key := this.InfoPrefix + uidStr + ":" + cidStr + ":" + idStr

	res, reserr := redisHandle.Do("hgetall", key)
	factoryData, err := redis.StringMap(res, reserr)

	if err != nil {
		return nil
	}

	factoryInfo = &common.Factory{}
	factoryInfo.Id, err = strconv.Atoi(factoryData["id"])
	if err != nil {
		return nil
	}

	factoryInfo.CollectTime, err = strconv.Atoi(factoryData["collectTime"])
	if err != nil {
		return nil
	}

	factoryInfo.Level, err = strconv.Atoi(factoryData["level"])
	if err != nil {
		return nil
	}

	return factoryInfo
}

func (this *Produce) GetAll(uid int, cid int) map[int]*common.Factory {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)

	key := "fac:" + cidStr + ":" + uidStr

	res, reserr := redisHandle.Do("hgetall", key)
	allData, err := redis.IntMap(res, reserr)

	if err != nil {
		return nil
	}

	factoryData := make(map[int]*common.Factory)
	for _, v := range allData {
		key := this.InfoPrefix + uidStr + ":" + cidStr + ":" + strconv.Itoa(v)
		res, reserr = redisHandle.Do("hgetall", key)
		info, err := redis.StringMap(res, reserr)

		if err != nil {
			continue
		}
		collectTime, err := strconv.Atoi(info["collectTime"])
		if err != nil {
			continue
		}
		level, err := strconv.Atoi(info["level"])
		if err != nil {
			continue
		}

		factoryData[v] = &common.Factory{Id: int(v), Level: level, CollectTime: collectTime}
	}

	return factoryData
}
