package pve

import (
	"slg/common"
	"slg/model"
	"strconv"
	"time"
)

func (this *Pve) get(param interface{}, extra interface{}) (output []byte, errno int) {

	if myData, ok := extra.(*common.UserData); ok {

		pveMapObj := model.NewPveMap()
		pveCityObj := model.NewPveCity()

		getMP := &GetStruct{}
		getMP.Id = CODE_GET
		getMP.Time = time.Now().Unix()

		getMP.MI = pveMapObj.GetAll(myData.Uid)
		citys := pveCityObj.GetCitys(myData.Uid)
		getMP.City = make(map[int]*CityInfo)
		for _, city := range citys {

			pveCityInfo := &CityInfo{}
			pveCityInfo.Cid, _ = strconv.Atoi(city["cid"])
			pveCityInfo.X, _ = strconv.Atoi(city["x"])
			pveCityInfo.Y, _ = strconv.Atoi(city["y"])
			pveCityInfo.Id, _ = strconv.Atoi(city["id"])

			getMP.City[pveCityInfo.Id] = pveCityInfo

		}

		serialize, err := common.Serialize(getMP)

		if err != nil {
			return []byte{}, common.CODE_ERR_UNKNOW
		}

		return serialize, 0

	}

	return []byte{}, 0
}
