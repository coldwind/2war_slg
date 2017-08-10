package hall

import (
	"time"
	"zjh/common"
)

func (this *Hall) get(param interface{}, extra interface{}) (output []byte, errno int) {
	if _, ok := param.(map[string]interface{}); !ok {
		return []byte{}, common.CODE_ERR_UNKNOW
	}

	if myData, ok := extra.(*common.UserData); ok {
		if !common.CheckLogin(myData) {
			return []byte{}, common.CODE_ERR_UNKNOW
		}

		hallListMP := &HallListMP{}
		hallListMP.Id = CODE_GET
		hallListMP.List = make([]*common.HallInfo, 0, 1024)
		hallListMP.Time = time.Now().Unix()

		for _, v := range common.HallData {
			hallListMP.List = append(hallListMP.List, v)
		}

		serialize, err := common.Serialize(hallListMP)
		if err == nil {
			return serialize, 0
		}
	}

	return []byte{}, common.CODE_ERR_UNKNOW
}
