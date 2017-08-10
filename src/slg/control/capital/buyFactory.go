package capital

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Capital) buyFactory(param interface{}, extra interface{}) (output []byte, errno int) {
	if paramData, ok := param.(map[string]interface{}); ok {

		if fid, ok := paramData["fid"].(uint64); ok {

			if myData, ok := extra.(*common.UserData); ok {

				pid := int(fid)
				cid64, _ := paramData["cid"].(uint64)
				cid := int(cid64)

				// 校验配置是否存在
				if _, ok := common.ProduceConfs[pid]; !ok {
					return []byte{}, common.CODE_ERR_UNKNOW
				}

				// 校验用户余额是否充足
				if myData.Gold < common.ProduceConfs[pid].Price {
					return []byte{}, common.CODE_ERR_UNKNOW
				}

				userObj := model.NewUser()

				// 校验用户是否已购买
				factoryObj := model.NewProduce()
				factory := factoryObj.Get(myData.Uid, cid, pid)
				newLevel := 1
				nowTime := int(time.Now().Unix())
				collectTime := nowTime
				gold := common.ProduceConfs[pid].Price
				if factory != nil {
					// 升级
					cosumeRes := userObj.IncrGold(myData.Uid, (0 - gold))
					if !cosumeRes {
						return []byte{}, common.CODE_ERR_UNKNOW
					}
					newLevel = factory.Level + 1

					factoryObj.SetLevel(myData.Uid, cid, pid, newLevel)
					collectTime = factory.CollectTime
				} else {
					// 购买
					cosumeRes := userObj.IncrGold(myData.Uid, (0 - gold))
					if !cosumeRes {
						return []byte{}, common.CODE_ERR_UNKNOW
					}
					factoryObj.Add(myData.Uid, cid, pid)
				}

				factoryMP := &buyStruct{Id: CODE_BUY_FACTORY, Fid: pid, Gold: gold, Level: newLevel, CollectTime: collectTime, Time: nowTime}
				serialize, err := common.Serialize(factoryMP)

				if err != nil {
					return []byte{}, common.CODE_ERR_UNKNOW
				}

				return serialize, 0
			}
		}
	}

	return []byte{}, common.CODE_ERR_UNKNOW

}
