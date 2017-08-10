package unlock

import (
	"log"
	"slg/common"
	"slg/model"
	"time"
)

func (this *Unlock) get(param interface{}, extra interface{}) (output []byte, errno int) {

	if paramData, ok := param.(map[string]interface{}); ok {

		if myData, ok := extra.(*common.UserData); ok {

			var cid int
			var unlockType int
			if _, ok := paramData["cid"].(uint64); !ok {
				cid = myData.MainCity
			} else {
				cid64, _ := paramData["cid"].(uint64)
				cid = int(cid64)
			}

			if _, ok := paramData["type"].(uint64); !ok {
				unlockType = model.CODE_UNLOCK_TYPE_MANAGE
			} else {
				unlockType64, _ := paramData["type"].(uint64)
				unlockType = int(unlockType64)
			}

			unlockObj := model.NewUnlock()

			getMP := &getStruct{}
			getMP.Id = CODE_GET
			getMP.Type = unlockType
			getMP.List = make([]int, 0)
			list := unlockObj.GetAll(myData.Uid, cid, unlockType)
			if len(list) > 0 {
				// TODO 过滤重复的上一级ID
				getMP.List = list
			}

			log.Println("list:", getMP.List)

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
