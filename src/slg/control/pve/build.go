package pve

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Pve) build(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			cid64, _ := paramData["cid"].(uint64)
			cid := int(cid64)

			x64, _ := paramData["x"].(uint64)
			x := int(x64)
			y64, _ := paramData["y"].(uint64)
			y := int(y64)

			pveCityObj := model.NewPveCity()

			cityId := pveCityObj.Add(myData.Uid, cid, x, y)
			if cityId == 0 {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			buildMP := &BuildStruct{}
			buildMP.Id = CODE_BUILD
			buildMP.Time = time.Now().Unix()
			buildMP.CityId = cityId

			// TODO 扣除统治力

			buildMP.Hp = 1

			serialize, err := common.Serialize(buildMP)

			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0

		}

	}
	return []byte{}, 0
}
