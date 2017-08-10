package army

import (
	"log"
	"slg/common"
	"slg/model"
	"time"
)

func (this *Army) makeArmy(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {
			var cid int
			var id int
			var num int
			var corpsId int
			if _, ok := paramData["cid"].(uint64); !ok {
				cid = myData.MainCity
			} else {
				cid64, _ := paramData["cid"].(uint64)
				cid = int(cid64)
			}

			if _, ok := paramData["uniId"].(uint64); !ok {
				log.Println("make error:1")
				return []byte{}, common.CODE_ERR_UNKNOW
			} else {
				id64 := paramData["uniId"].(uint64)
				id = int(id64)
			}

			if _, ok := paramData["num"].(uint64); !ok {
				log.Println("make error:2")

				return []byte{}, common.CODE_ERR_UNKNOW
			} else {
				num64 := paramData["num"].(uint64)
				num = int(num64)
			}

			if _, ok := paramData["corpsId"].(uint64); !ok {
				log.Println("make error:3")

				return []byte{}, common.CODE_ERR_UNKNOW
			} else {
				corpsId64, _ := paramData["corpsId"].(uint64)
				corpsId = int(corpsId64)
			}

			armyObj := model.NewArmy()
			unlockObj := model.NewUnlock()
			// 校验用户是否有资格生产
			if !unlockObj.ArmyIsUnlock(myData.Uid, cid, id) {
				log.Println("make error:4")

				return []byte{}, common.CODE_ERR_UNKNOW
			}

			// 校验用户是否有足够的金币
			price := common.ArmyConfs[id].MakePrice * num
			if price > myData.Gold {
				log.Println("make error:5")

				return []byte{}, common.CODE_ERR_UNKNOW
			}

			oil := common.ArmyConfs[id].MakeOil * num

			userObj := model.NewUser()

			// 扣钱
			if price > 0 {
				cosumeRes := userObj.IncrGold(myData.Uid, (0 - price))
				if !cosumeRes {
					log.Println("make error:6")

					return []byte{}, common.CODE_ERR_UNKNOW
				}
			}

			// 扣油
			if oil > 0 {
				cosumeRes := userObj.IncrOil(myData.Uid, (0 - oil))
				if !cosumeRes {
					log.Println("make error:7")

					return []byte{}, common.CODE_ERR_UNKNOW
				}
			}

			// 生产
			aid := common.ArmyConfs[id].ArmyId
			makeRes := armyObj.AddArmy(myData.Uid, cid, corpsId, aid, num)
			if !makeRes {
				log.Println("make error:8")

				return []byte{}, common.CODE_ERR_UNKNOW
			}

			makeMP := &makeStruct{}
			makeMP.Id = CODE_MAKE_ARMY
			makeMP.Oil = oil
			makeMP.Gold = price
			makeMP.Time = time.Now().Unix()

			serialize, err := common.Serialize(makeMP)

			if err != nil {
				log.Println("make error:9")

				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0
		}
	}

	return []byte{}, 0
}
