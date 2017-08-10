package army

import (
	"slg/common"
	"slg/model"
	"time"
)

func (this *Army) getCorps(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			var cid int
			if _, ok := paramData["cid"].(uint64); !ok {
				cid = myData.MainCity
			} else {
				cid64, _ := paramData["cid"].(uint64)
				cid = int(cid64)
			}

			unlockObj := model.NewUnlock()

			getMP := &getCorpsStruct{}
			getMP.Id = CODE_GET_CORPS

			getMP.List = unlockObj.GetAll(myData.Uid, cid, model.CODE_UNLOCK_TYPE_CORPS)

			getMP.Time = time.Now().Unix()

			serialize, err := common.Serialize(getMP)

			if err != nil {
				return []byte{}, common.CODE_ERR_UNKNOW
			}

			return serialize, 0

		}
	}

	return []byte{}, 0
}
