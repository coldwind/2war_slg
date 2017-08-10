package model

import (
	"errors"
	"log"
	"slg/common"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type User struct {
	UserBasePrefix string
}

func NewUser() (user *User) {
	user = &User{}
	user.UserBasePrefix = "ub:"

	return
}

func (this *User) GetUserInfo(uid int) (userInfo map[string]string, err error) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	key := "ub:" + strconv.Itoa(uid)
	res, reserr := redisHandle.Do("hgetall", key)
	if reserr != nil {
		log.Println(reserr)
	}
	userBase, err := redis.StringMap(res, reserr)

	if err != nil {
		return map[string]string{}, errors.New("empty userinfo")
	}

	return userBase, nil
}

func (this *User) IncrGold(uid int, gold int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	key := "ub:" + strconv.Itoa(uid)
	_, err := redisHandle.Do("hincrby", key, "gold", gold)
	if err != nil {
		log.Println(err)
		return false
	}

	// 减mydata数据(gold为负)
	common.OnlineUser[uid].Gold += gold

	return true
}

func (this *User) IncrOil(uid int, oil int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	key := "ub:" + strconv.Itoa(uid)
	_, err := redisHandle.Do("hincrby", key, "oil", oil)
	if err != nil {
		log.Println(err)
		return false
	}

	// 减mydata数据(gold为负)
	common.OnlineUser[uid].Oil += oil

	return true
}

func (this *User) InitUser(uid int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	// 初始化基础信息
	key := "ub:" + strconv.Itoa(uid)

	res, reserr := redisHandle.Do("hmset", key,
		"gold", common.InitUserConf.Gold,
		"level", common.InitUserConf.Level,
		"exp", common.InitUserConf.Exp,
		"hp", common.InitUserConf.Hp,
		"oil", common.InitUserConf.Oil)
	if reserr != nil {
		log.Println(reserr)
	}
	_, err := redis.Int(res, reserr)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (this *User) SetMainCity(uid int, cityId int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	// 初始化基础信息
	key := "ub:" + strconv.Itoa(uid)
	cityIdStr := strconv.Itoa(cityId)
	res, reserr := redisHandle.Do("hset", key, "mainCity", cityIdStr)
	if reserr != nil {
		log.Println(reserr)
	}

	_, err := redis.Int(res, reserr)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
