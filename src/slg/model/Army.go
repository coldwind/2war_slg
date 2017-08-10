package model

import (
	"slg/common"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type Army struct {
	ArmyPrefix string // 军队前缀
}

func NewArmy() (armyObj *Army) {
	armyObj = &Army{}
	armyObj.ArmyPrefix = "army:"

	return
}

func (this *Army) GetArmyByCid(uid int, cid int, corpsId int) (army map[int]int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	corpsIdStr := strconv.Itoa(corpsId)

	key := this.ArmyPrefix + uidStr + ":" + cidStr + ":" + corpsIdStr
	army = make(map[int]int)

	res, reserr := redisHandle.Do("hgetall", key)
	if reserr != nil {
		return army
	}

	var kInt int
	armyRes, _ := redis.IntMap(res, reserr)
	for k, v := range armyRes {
		kInt, _ = strconv.Atoi(k)
		army[kInt] = v
	}

	return
}

func (this *Army) AddArmy(uid int, cid int, corpsId int, aid int, num int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	corpsIdStr := strconv.Itoa(corpsId)

	key := this.ArmyPrefix + uidStr + ":" + cidStr + ":" + corpsIdStr

	res, reserr := redisHandle.Do("hincrby", key, aid, num)

	if reserr != nil {
		return false
	}

	result, err := redis.Bool(res, reserr)
	if err != nil {
		return false
	}

	return result
}

func (this *Army) SubtractArmy(uid int, cid int, corpsId int, aid int, num int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	corpsIdStr := strconv.Itoa(corpsId)

	key := this.ArmyPrefix + uidStr + ":" + cidStr + ":" + corpsIdStr

	res, reserr := redisHandle.Do("hincrby", key, aid, (0 - num))

	result, err := redis.Bool(res, reserr)
	if err != nil {
		return false
	}

	return result
}
