package unlock

import (
	"log"
	"slg/common"
	"slg/model"
	"time"
)

func (this *Unlock) unlock(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			var cid int
			var mid int
			var unlockType int

			if _, ok := paramData["cid"].(uint64); !ok {
				cid = myData.MainCity
			} else {
				cid64, _ := paramData["cid"].(uint64)
				cid = int(cid64)
			}
			if _, ok := paramData["mid"].(uint64); !ok {
				log.Println("mid empty")
				return []byte{}, common.CODE_ERR_UNKNOW
			} else {
				mid64, _ := paramData["mid"].(uint64)
				mid = int(mid64)
			}

			if _, ok := paramData["type"].(uint64); !ok {
				unlockType = model.CODE_UNLOCK_TYPE_MANAGE
			} else {
				unlockType64, _ := paramData["type"].(uint64)
				unlockType = int(unlockType64)
			}

			var prevId int
			var price int

			// 获取配置
			if unlockType == model.CODE_UNLOCK_TYPE_MANAGE {
				if _, ok := common.UnlockManageConfs[mid]; !ok {
					return []byte{}, common.CODE_ERR_UNKNOW
				}

				prevId = common.UnlockManageConfs[mid].PrevId
				price = common.UnlockManageConfs[mid].UnlockPrice

			} else if unlockType == model.CODE_UNLOCK_TYPE_ARMY {
				if _, ok := common.ArmyConfs[mid]; !ok {
					return []byte{}, common.CODE_ERR_UNKNOW
				}

				prevId = common.ArmyConfs[mid].PrevId
				price = common.ArmyConfs[mid].UnlockPrice
			} else if unlockType == model.CODE_UNLOCK_TYPE_CORPS {
				if _, ok := common.UnlockCorpsConfs[mid]; !ok {
					return []byte{}, common.CODE_ERR_UNKNOW
				}

				prevId = common.UnlockCorpsConfs[mid].PrevId
				price = common.UnlockCorpsConfs[mid].UnlockPrice
			} else {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			unlockObj := model.NewUnlock()

			// 校验是否有前置条件
			if prevId > 0 {
				// 校验前置ID是否已解锁
				isUnlock := unlockObj.IsUnlock(myData.Uid, cid, prevId, unlockType)
				if !isUnlock {
					log.Println("prev unlock", prevId)
					return []byte{}, common.CODE_ERR_UNKNOW
				}
			}

			// 判断用户金币是否足够
			if myData.Gold < price {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			// 扣钱
			userObj := model.NewUser()
			cosumeRes := userObj.IncrGold(myData.Uid, (0 - price))
			if !cosumeRes {
				log.Println(cosumeRes)
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			// 解锁新ID
			unlockRes := unlockObj.Unlock(myData.Uid, cid, mid, unlockType)
			if !unlockRes {
				log.Println(unlockRes)
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			// 删除上一次ID
			if prevId > 0 {
				unlockObj.Remove(myData.Uid, cid, prevId, unlockType)
			}

			unlockMP := &unlockStruct{}
			unlockMP.Id = CODE_UNLOCK
			unlockMP.Type = unlockType
			unlockMP.Gold = price
			unlockMP.Mid = mid
			unlockMP.Time = time.Now().Unix()

			serialize, err := common.Serialize(unlockMP)

			if err != nil {
				log.Println(err)
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0
		}
	}

	return []byte{}, 0
}
