package common

const TRANS_MAX_LEN uint16 = 1024
const REDIS_POOL_LEN int = 200

// 网络数据缓存模块
type NetWorkCache struct {
	Len   int
	Cache []byte
}

var RunLogHandle *LogSys
var DbHandle *Dao
var RedisHandle *RedisDao
var OnlineUser map[int]*UserData
var HallData map[int]*HallInfo
var HallCount int
