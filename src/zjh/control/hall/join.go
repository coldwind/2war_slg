package hall

import (
	"time"
	"zjh/common"
)

func (this *Hall) join(param interface{}, extra interface{}) (output []byte, errno int) {
	if joinData, ok := param.(map[string]interface{}); ok {

		return []byte{}, common.CODE_ERR_UNKNOW

		if myData, ok := extra.(*common.UserData); ok {
			if !common.CheckLogin(myData) {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			hallId := int(joinData["hallId"].(uint64))

			common.HallData[hallId].Users = append(common.HallData[hallId].Users, myData)

			// 推送给其他用户加入信息

			joinMP := &JoinMP{}
			joinMP.Id = CODE_JOIN
			joinMP.HallInfo = common.HallData[hallId]
			joinMP.Time = time.Now().Unix()

			serialize, err := common.Serialize(joinMP)
			if err == nil {
				return serialize, 0
			}
		}
	}

	return []byte{}, common.CODE_ERR_UNKNOW
}
