package capital

import (
	"log"
	"slg/common"
	"slg/model"
	"time"
)

func (this *Capital) AutoCollect(uid int) {
	if _, ok := common.OnlineUser[uid]; ok {
		// 获取用户CITY
		pveCityObj := model.NewPveCity()
		citys := pveCityObj.GetCityIds(uid)

		unlockObj := model.NewUnlock()
		factoryObj := model.NewProduce()

		nowTime := int(time.Now().Unix())
		var collectTime int
		income := 0

		// 循环用户CITY获取MANAGE
		for _, cid := range citys {

			log.Println("城市：", cid)

			// 计算收益&收集
			mids := unlockObj.GetAll(uid, cid, model.CODE_UNLOCK_TYPE_MANAGE)
			for _, mid := range mids {
				log.Println("经理:", mid)
				// 校验配置文件
				if mangeConf, ok := common.UnlockManageConfs[mid]; ok {
					// 获取factory内容
					factoryInfo := factoryObj.Get(uid, cid, mangeConf.FactoryId)
					if factoryInfo == nil {
						continue
					}

					// 判断收取条件
					if factoryInfo.CollectTime > nowTime {
						continue
					}

					// 收取--修改时间 & 加钱
					collectTime = factoryInfo.CollectTime + common.ProduceConfs[mangeConf.FactoryId].Interval
					factoryObj.Set(uid, cid, mangeConf.FactoryId, collectTime)
					income += common.ProduceConfs[mangeConf.FactoryId].Produce
				}
				log.Println("收入:", income)
			}
		}

		// 计算收益 推送给用户
		if income > 0 {
			userObj := model.NewUser()
			userObj.IncrGold(uid, income)

			pushIncomeMP := &common.PushIncomeStruct{}
			pushIncomeMP.Id = common.PUSH_INCOME_CODE
			pushIncomeMP.Income = income

			serialize, err := common.Serialize(pushIncomeMP)

			if err == nil {
				common.Response(*common.OnlineUser[uid].Conn, serialize)
			}
		}
	}
}
