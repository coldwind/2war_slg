package pve

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Pve) dismantle(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			cityId64, _ := paramData["cityId"].(uint64)
			cityId := int(cityId64)

			pveCityObj := model.NewPveCity()

			cityId = pveCityObj.Del(myData.Uid, cityId)

			if cityId == 0 {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			dismantleMP := &DismantleStruct{}
			dismantleMP.Id = CODE_DISMANTLE
			dismantleMP.Time = time.Now().Unix()
			dismantleMP.CityId = cityId

			serialize, err := common.Serialize(dismantleMP)

			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0

		}

	}
	return []byte{}, 0
}
