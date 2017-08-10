package capital

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Capital) get(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			var cid int
			if _, ok := paramData["cid"].(uint64); !ok {
				cid = myData.MainCity
			} else {
				cid64, _ := paramData["cid"].(uint64)
				cid = int(cid64)
			}
			factoryObj := model.NewProduce()
			factorys := factoryObj.GetAll(myData.Uid, cid)

			getMP := &getStruct{}
			getMP.Id = CODE_GET
			getMP.List = make([]*FactoryStruct, 0)
			getMP.Time = time.Now().Unix()

			for _, v := range factorys {
				factoryInfo := &FactoryStruct{Id: v.Id, Level: v.Level, CollectTime: v.CollectTime}
				getMP.List = append(getMP.List, factoryInfo)
			}

			serialize, err := common.Serialize(getMP)

			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0

		}
	}

	return []byte{}, 0
}
