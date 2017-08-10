package model

import (
	"errors"
	"log"
	"strconv"
	"zjh/common"

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

func (this *User) InitUser(uid int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	// 初始化基础信息
	key := "ub:" + strconv.Itoa(uid)
	log.Println("redis key:", key)

	res, reserr := redisHandle.Do("hmset", key,
		"gold", "100",
		"level", "1",
		"exp", "0",
		"hp", "10")
	if reserr != nil {
		log.Println("redis set error:", reserr)
	}
	_, err := redis.String(res, reserr)

	if err != nil {
		log.Println("redis get error", err)
		return false
	}

	return true
}
