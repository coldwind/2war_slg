package model

import (
	"log"
	"slg/common"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type Unlock struct {
	UnlockManagePrefix string // 经理人解锁前缀
	UnlockArmyPrefix   string // 解锁军队前缀
	UnlockCorpsPrefix  string // 解锁军团前缀

}

func NewUnlock() (unlockObj *Unlock) {
	unlockObj = &Unlock{}
	unlockObj.UnlockManagePrefix = "unlock:manage:"
	unlockObj.UnlockArmyPrefix = "unlock:army:"
	unlockObj.UnlockCorpsPrefix = "unlock:corps:"

	return
}

func (this *Unlock) GetAll(uid int, cid int, lockType int) (mIds []int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	mIds = make([]int, 0, 10)
	var key string

	if lockType == CODE_UNLOCK_TYPE_MANAGE {
		key = this.UnlockManagePrefix + uidStr + ":" + cidStr

	} else if lockType == CODE_UNLOCK_TYPE_ARMY {
		key = this.UnlockArmyPrefix + uidStr + ":" + cidStr
	} else if lockType == CODE_UNLOCK_TYPE_CORPS {
		key = this.UnlockCorpsPrefix + uidStr + ":" + cidStr
	}
	res, reserr := redisHandle.Do("smembers", key)
	if reserr != nil {
		log.Println(reserr)
		return mIds
	}

	mIds, _ = redis.Ints(res, reserr)

	return mIds
}

func (this *Unlock) Unlock(uid int, cid int, mid int, lockType int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	var key string

	if lockType == CODE_UNLOCK_TYPE_MANAGE {
		key = this.UnlockManagePrefix + uidStr + ":" + cidStr

	} else if lockType == CODE_UNLOCK_TYPE_ARMY {
		key = this.UnlockArmyPrefix + uidStr + ":" + cidStr
	} else if lockType == CODE_UNLOCK_TYPE_CORPS {
		key = this.UnlockCorpsPrefix + uidStr + ":" + cidStr
	}

	_, reserr := redisHandle.Do("sadd", key, mid)
	if reserr != nil {
		log.Println(reserr)

		return false
	}

	return true
}

func (this *Unlock) IsUnlock(uid int, cid int, mid int, lockType int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	var key string

	if lockType == CODE_UNLOCK_TYPE_MANAGE {
		key = this.UnlockManagePrefix + uidStr + ":" + cidStr

	} else if lockType == CODE_UNLOCK_TYPE_ARMY {
		key = this.UnlockArmyPrefix + uidStr + ":" + cidStr
	} else if lockType == CODE_UNLOCK_TYPE_CORPS {
		key = this.UnlockCorpsPrefix + uidStr + ":" + cidStr
	}

	res, reserr := redisHandle.Do("sismember", key, mid)
	if reserr != nil {
		log.Println(reserr)

		return false
	}

	isUnlock, err := redis.Bool(res, reserr)
	if err != nil {
		log.Println(err)

		return false
	}

	return isUnlock
}

func (this *Unlock) Remove(uid int, cid int, mid int, lockType int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	var key string

	if lockType == CODE_UNLOCK_TYPE_MANAGE {
		key = this.UnlockManagePrefix + uidStr + ":" + cidStr

	} else if lockType == CODE_UNLOCK_TYPE_ARMY {
		key = this.UnlockArmyPrefix + uidStr + ":" + cidStr
	} else if lockType == CODE_UNLOCK_TYPE_CORPS {
		key = this.UnlockCorpsPrefix + uidStr + ":" + cidStr
	}
	_, reserr := redisHandle.Do("srem", key, mid)
	if reserr != nil {
		log.Println(reserr)
		return false
	}

	return true
}

func (this *Unlock) GetAllArmy(uid int, cid int) (mIds []int) {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)
	mIds = make([]int, 0, 10)

	key := this.UnlockArmyPrefix + uidStr + ":" + cidStr
	res, reserr := redisHandle.Do("smembers", key)
	if reserr != nil {
		log.Println(reserr)
		return mIds
	}

	mIds, _ = redis.Ints(res, reserr)

	return mIds
}

func (this *Unlock) ArmyIsUnlock(uid int, cid int, mid int) bool {

	redisHandle := common.RedisHandle.GetRedisHandle(uid)
	defer redisHandle.Close()

	uidStr := strconv.Itoa(uid)
	cidStr := strconv.Itoa(cid)

	key := this.UnlockArmyPrefix + uidStr + ":" + cidStr

	res, reserr := redisHandle.Do("sismember", key, mid)
	if reserr != nil {
		return false
	}

	isUnlock, err := redis.Bool(res, reserr)
	if err != nil {
		return false
	}

	return isUnlock
}
