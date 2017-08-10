package capital

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Capital) collect(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if fid, ok := paramData["fid"].(uint64); ok {

			if myData, ok := extra.(*common.UserData); ok {

				pid := int(fid)
				cid64, _ := paramData["cid"].(uint64)
				cid := int(cid64)

				factoryObj := model.NewProduce()
				factory := factoryObj.Get(myData.Uid, cid, pid)

				// 判断是否可以收取
				if factory == nil {
					return []byte{}, 0
				}

				// 计算收取金币数量
				nowTime := int(time.Now().Unix())
				nextCollectTime := nowTime + common.ProduceConfs[pid].Interval
				if factory.CollectTime > nowTime {
					return []byte{}, 0
				}

				// 收取
				res := factoryObj.Set(myData.Uid, cid, pid, nextCollectTime)
				if res {
					// 加钱
					userObj := model.NewUser()
					userObj.IncrGold(myData.Uid, common.ProduceConfs[pid].Produce)

					collectMP := &collectStruct{}
					collectMP.Id = CODE_COLLECT
					collectMP.Gold = common.ProduceConfs[pid].Produce
					collectMP.Fid = pid
					collectMP.Time = nowTime
					collectMP.CollectTime = nextCollectTime
					serialize, err := common.Serialize(collectMP)
					if err != nil {
						return []byte{}, common.CODE_ERR_UNKNOW
					}

					return serialize, 0
				}

			}
		}
	}

	return []byte{}, common.CODE_ERR_UNKNOW
}
