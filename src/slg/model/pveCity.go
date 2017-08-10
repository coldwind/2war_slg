package model

import (
	"slg/common"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type PveCity struct {
	CityVectorPrefix string // 城市容器
	CityIdPrefix     string // 城市自增ID
	CityInfoPrefix   string // 城市信息
}

func NewPveCity() (pveCity *PveCity) {
	pveCity = &PveCity{}
	pveCity.CityIdPrefix = "pve:cid:"
	pveCity.CityVectorPrefix = "pve:cv:"
	pveCity.CityInfoPrefix = "pve:ci:"

	return pveCity
}

func (this *PveCity) Add(uid int, cid int, x int, y int) (cityId int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)

	// 自增ID
	key := this.CityIdPrefix + uidStr
	res, reserr := redisHandle.Do("incrby", key, "1")
	if reserr != nil {
		return 0
	}

	id, err := redis.Int(res, reserr)
	if err != nil {
		return 0
	}

	// 写入城市容器
	key = this.CityVectorPrefix + uidStr
	_, reserr = redisHandle.Do("sadd", key, id)
	if reserr != nil {
		return 0
	}

	// 写入城市信息
	idStr := strconv.Itoa(id)
	cidStr := strconv.Itoa(cid)
	key = this.CityInfoPrefix + uidStr + ":" + idStr
	_, reserr = redisHandle.Do("hmset", key, "id", idStr, "cid", cidStr, "level", "1", "x", x, "y", y)
	if reserr != nil {
		return 0
	}

	return id
}

func (this *PveCity) GetCityIds(uid int) (cityIds []int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	key := this.CityVectorPrefix + uidStr

	cityIds = make([]int, 0, 10)

	res, reserr := redisHandle.Do("smembers", key)
	if reserr != nil {
		return cityIds
	}

	cityIds, _ = redis.Ints(res, reserr)

	return cityIds
}

func (this *PveCity) GetCitys(uid int) (citys []map[string]string) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	key := this.CityVectorPrefix + uidStr

	cityIds := make([]int, 0, 10)
	citys = make([]map[string]string, 0, 10)

	res, reserr := redisHandle.Do("smembers", key)
	if reserr != nil {
		return citys
	}

	cityIds, _ = redis.Ints(res, reserr)

	for _, v := range cityIds {
		key = this.CityInfoPrefix + uidStr + ":" + strconv.Itoa(v)
		res, reserr := redisHandle.Do("hgetall", key)
		if reserr != nil {
			continue
		}
		cityData, err := redis.StringMap(res, reserr)
		if err == nil {
			citys = append(citys, cityData)
		}
	}

	return citys
}

func (this *PveCity) Del(uid int, id int) (cityId int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)

	// 删除城市容器
	key := this.CityVectorPrefix + uidStr
	_, reserr := redisHandle.Do("srem", key, id)
	if reserr != nil {
		return 0
	}

	// 删除城市信息
	idStr := strconv.Itoa(id)
	key = this.CityInfoPrefix + uidStr + ":" + idStr
	_, reserr = redisHandle.Do("del", key)
	if reserr != nil {
		return 0
	}

	return id
}
