package account

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"zjh/common"
	"zjh/model"
)

func (this *Account) reg(param interface{}, extra interface{}) (output []byte, errno int) {

	if _, ok := param.(map[string]interface{}); !ok {
		return []byte{}, common.CODE_ERR_UNKNOW
	}

	mapData := param.(map[string]interface{})

	if _, ok := mapData["name"]; !ok {
		return []byte{}, CODE_ERR_USER_NOT_EXIST
	}

	if _, ok := mapData["password"]; !ok {
		return []byte{}, CODE_ERR_USER_NOT_EXIST
	}

	user := mapData["name"].(string)
	password := mapData["password"].(string)

	where := fmt.Sprintf("username='%s'", user)
	userInfo := common.DbHandle.GetRecord("user", where, 0)

	if userInfo != nil {
		return []byte{}, CODE_ERR_USER_EXIST
	}

	password = common.Md5Encode(password)

	// 写入用户表
	var fields []string
	var values []string
	nowtime := int(time.Now().Unix())
	regtime := strconv.Itoa(nowtime)

	fields = append(fields, "username")
	fields = append(fields, "password")
	fields = append(fields, "regtime")
	fields = append(fields, "logintime")

	values = append(values, user)
	values = append(values, password)
	values = append(values, regtime)
	values = append(values, regtime)

	res := common.DbHandle.AddRecord("user", fields, values, 0)

	if res {
		userInfo = common.DbHandle.GetRecord("user", where, 0)
		uid, err := strconv.Atoi(userInfo["id"])
		if err != nil {
			return []byte{}, common.CODE_ERR_UNKNOW
		}

		// 初始化用户信息
		userObj := &model.User{}
		userObj.InitUser(uid)

		regObj := &RegMPStruct{}
		regObj.Id = CODE_REG
		regObj.Time = time.Now().Unix()
		log.Println(regObj)
		regMP, err := common.Serialize(regObj)

		if err != nil {
			return []byte{}, common.CODE_ERR_UNKNOW
		}

		return regMP, 0
	}

	return []byte{}, common.CODE_ERR_UNKNOW
}
