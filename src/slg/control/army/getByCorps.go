package army

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Army) getByCorps(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			var cid int
			var corpsId int
			if _, ok := paramData["cid"].(uint64); !ok {
				cid = myData.MainCity
			} else {
				cid64, _ := paramData["cid"].(uint64)
				cid = int(cid64)
			}

			if _, ok := paramData["corpsId"].(uint64); !ok {
				return []byte{}, common.CODE_ERR_UNKNOW
			} else {
				corpsId64, _ := paramData["corpsId"].(uint64)
				corpsId = int(corpsId64)
			}

			armyObj := model.NewArmy()
			unlockObj := model.NewUnlock()

			getMP := &getArmyStruct{}
			getMP.Id = CODE_GET_ARMY_CORPS
			getMP.CorpsId = corpsId
			getMP.List = make(map[int]*armyStruct)

			armyUnlockIds := unlockObj.GetAllArmy(myData.Uid, cid)

			armyData := armyObj.GetArmyByCid(myData.Uid, cid, corpsId)
			// 通过lock id计算出士兵的信息
			for _, unlockId := range armyUnlockIds {

				armyInfo := &armyStruct{}
				armyInfo.Id = unlockId
				armyInfo.Aid = common.ArmyConfs[unlockId].ArmyId
				armyInfo.Num = 0

				for k, v := range armyData {
					if common.ArmyConfs[unlockId].ArmyId == k {
						armyInfo.Num = v
					}
				}

				getMP.List[armyInfo.Aid] = armyInfo

			}

			getMP.Time = time.Now().Unix()

			serialize, err := common.Serialize(getMP)

			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0

		}
	}

	return []byte{}, 0
}
