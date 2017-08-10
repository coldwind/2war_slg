package network

import (
	"encoding/json"
	"log"
	"math"
	"time"
	"zjh/common"
	"zjh/control/account"
	"zjh/control/hall"

	"github.com/gorilla/websocket"
)

type Route struct {
}

func (this *Route) Run(conn *websocket.Conn) {
	// 初始化用户基础数据结构
	myData := &common.UserData{}
	myData.RefreshTime = time.Now().Unix()
	myData.Conn = conn
	myData.Uid = 0

	defer destructUser(myData)

	// 请求的接口标识
	var flag int = -1
	var id int = 0

	// 定义解析JSON的变量
	//	var parseData = make(map[string]interface{})
	//	var err error

	accountObj := &account.Account{}
	hallObj := &hall.Hall{}
	//	errorCount := 0

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	parseData := make(map[string]interface{}, 0)
	var outputData []byte

	for {

		defer conn.Close()
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		err = json.Unmarshal(message, &parseData)

		if err != nil {
			continue
		}

		log.Println(parseData)

		if pv, ok := parseData["id"].(uint64); ok {
			flag = int(math.Floor(float64(pv / 100)))
			id = int(pv)
		} else if pv, ok := parseData["id"].(float64); ok {
			flag = int(math.Floor(float64(pv / 100)))
			id = int(pv)
		} else {
			conn.Close()
			goto Exit
		}

		switch flag {
		case CODE_ROUTE_PING:
			// 刷新心跳时间
			outputData = formatPong(id)

		case CODE_ROUTE_ACCOUNT:
			// 用户模块
			outputData = accountObj.Init(id, parseData, myData)

		case CODE_ROUTE_HALL:
			// 大厅模块
			outputData = hallObj.Init(id, parseData, myData)

		case 999:
			outputData = []byte{}
			// 退出游戏
			conn.Close()
			goto Exit
		default:
			outputData = []byte{}

			common.Response(conn, common.GetErrFormat(common.CODE_ERR_UNKNOW))
			conn.Close()
			goto Exit
		}

		//		if err != nil {
		//			fmt.Println("error", err)
		//			conn.Close()
		//			goto Exit
		//		}
		//		errorCount = 0
		//		fmt.Println("output:", outputData)
		common.Response(conn, outputData)
	}
Exit:
}

func destructUser(myData *common.UserData) {
	if myData.Uid > 0 {
		// 清除用户集数据
		delete(common.OnlineUser, myData.Uid)
	}
}
