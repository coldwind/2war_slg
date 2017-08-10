package account

import (
	"fmt"
	"log"
	"slg/common"
	"slg/model"
	"strconv"
	"time"
)

func (this *Account) login(param interface{}, extra interface{}) (output []byte, errno int) {

	log.Println("login")

	if _, ok := param.(map[string]interface{}); !ok {
		log.Println("param error:", ok)
		return []byte{}, common.CODE_ERR_UNKNOW
	}

	mapData := param.(map[string]interface{})
	log.Println("mapdata:", mapData)

	if myData, ok := extra.(*common.UserData); ok {
		log.Println("mydata:", myData)

		if _, ok := mapData["name"]; !ok {
			log.Println("mapdata name error:", ok)

			return []byte{}, CODE_ERR_USER_NOT_EXIST
		}

		if _, ok := mapData["password"]; !ok {
			log.Println("mapdata password error:", ok)
			return []byte{}, CODE_ERR_USER_NOT_EXIST
		}

		user := mapData["name"].(string)
		password := mapData["password"].(string)

		where := fmt.Sprintf("username='%s'", user)
		userInfo := common.DbHandle.GetRecord("user", where, 0)

		if userInfo == nil {
			log.Println("userInfo error:")
			return []byte{}, CODE_ERR_USER_NOT_EXIST
		}

		encodePwd := common.Md5Encode(password)
		userPwd := userInfo["password"]
		if encodePwd != userPwd {
			log.Println("password error:", encodePwd, userPwd)

			return []byte{}, CODE_ERR_PASSWORD
		}

		log.Println("user info:", userInfo)
		myUid, err := strconv.Atoi(userInfo["id"])
		if err != nil {
			return []byte{}, CODE_ERR_USER_NOT_EXIST
		}

		myData.Uid = myUid

		myData.Nickname = userInfo["username"]

		loginMP := &LoginMPStruct{}
		loginMP.Id = CODE_LOGIN
		loginMP.Uid = myData.Uid
		loginMP.Time = time.Now().Unix()

		// 获取城市ID列表
		pveCityObj := model.NewPveCity()
		loginMP.City = pveCityObj.GetCityIds(myData.Uid)

		userObj := &model.User{}
		log.Println("myUid:", myUid)
		userBase, err := userObj.GetUserInfo(myUid)
		log.Println("user base:", userBase)

		if err != nil {
			log.Println("login error")

			return []byte{}, CODE_ERR_USER_NOT_EXIST
		}

		if _, ok = userBase["gold"]; ok {
			loginMP.Gold, err = strconv.Atoi(userBase["gold"])
			if err != nil {
				loginMP.Gold = 0
			}
		} else {
			loginMP.Gold = 0
		}

		if _, ok = userBase["hp"]; ok {
			loginMP.Hp, err = strconv.Atoi(userBase["hp"])
			if err != nil {
				loginMP.Hp = 0
			}
		} else {
			loginMP.Hp = 0
		}

		if _, ok = userBase["level"]; ok {
			loginMP.Level, err = strconv.Atoi(userBase["level"])
			if err != nil {
				loginMP.Level = 1
			}
		} else {
			loginMP.Level = 1
		}

		if _, ok = userBase["exp"]; ok {
			loginMP.Exp, err = strconv.Atoi(userBase["exp"])
			if err != nil {
				loginMP.Exp = 0
			}
		} else {
			loginMP.Exp = 0
		}

		if _, ok = userBase["mainCity"]; ok {
			loginMP.MainCity, err = strconv.Atoi(userBase["mainCity"])
		}

		myData.Level = loginMP.Level
		myData.Exp = loginMP.Exp
		myData.Hp = loginMP.Hp
		myData.Gold = loginMP.Gold
		myData.MainCity = loginMP.MainCity

		// 写入全局变量
		common.OnlineUser[myData.Uid] = myData

		loginMP.Nickname = userInfo["username"]

		serialize, err := common.Serialize(loginMP)

		return serialize, 0
	}

	log.Println("login error")

	return []byte{}, common.CODE_ERR_UNKNOW
}
