package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConf struct {
	Ip       string "ip"
	Port     string "port"
	Password string "password"
}

type RedisDao struct {
	TotalDB     int
	RedisHandle []*redis.Pool
}

// 获取当前key对应的服务器配置
func (this *RedisDao) GetRedisHandle(key int) (handle redis.Conn) {
	redisKey := key % this.TotalDB
	return this.RedisHandle[redisKey].Get()
}

// 初始化数据库
func (this *RedisDao) InitConf() {
	ConfString, err := ioutil.ReadFile("conf/redis.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var redisConf []RedisConf
	err = json.Unmarshal(ConfString, &redisConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range redisConf {

		redisClient := &redis.Pool{
			// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
			MaxIdle:     128,
			MaxActive:   REDIS_POOL_LEN,
			IdleTimeout: 300 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", v.Ip+":"+v.Port)
				if err != nil {
					return nil, err
				}
				// 选择db
				//c.Do("SELECT", REDIS_DB)
				return c, nil
			},
		}

		this.RedisHandle = append(this.RedisHandle, redisClient)
	}

	this.TotalDB = len(this.RedisHandle)
}
