package hall

import (
	"time"
	"zjh/common"
)

func (this *Hall) create(param interface{}, extra interface{}) (output []byte, errno int) {
	if _, ok := param.(map[string]interface{}); !ok {
		return []byte{}, common.CODE_ERR_UNKNOW
	}

	if myData, ok := extra.(*common.UserData); ok {
		if !common.CheckLogin(myData) {
			return []byte{}, common.CODE_ERR_UNKNOW
		}

		common.HallData[common.HallCount] = &common.HallInfo{}
		common.HallData[common.HallCount].HallId = common.HallCount
		common.HallData[common.HallCount].Name = "normal hall"
		common.HallData[common.HallCount].Uid = myData.Uid
		common.HallData[common.HallCount].Users = make([]*common.UserData, 0, 10)
		common.HallData[common.HallCount].Users = append(common.HallData[common.HallCount].Users, myData)

		createHallMP := &CreateHallMP{}
		createHallMP.Id = CODE_CREATE
		createHallMP.HallId = common.HallCount
		createHallMP.Time = time.Now().Unix()
		common.HallCount++

		serialize, err := common.Serialize(createHallMP)
		if err == nil {
			return serialize, 0
		}
	}

	return []byte{}, common.CODE_ERR_UNKNOW
}
