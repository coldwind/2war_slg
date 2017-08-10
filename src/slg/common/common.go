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
var ProduceConfs map[int]*ProduceStruct
var ArmyConfs map[int]*ArmyStruct
var GroundTypeConfs map[int]*GroundTypeStruct
var PveMapConfs map[int]*PveMapStruct
var InitUserConf *InitUserStruct
var InitPveMapConf []int
var UnlockManageConfs map[int]*UnlockManageStruct
var UnlockCorpsConfs map[int]*UnlockCorpsStruct
var NpcCorpsConfs []*NpcCorps
