package network

import (
	"fmt"
	"math"
	"net"
	"slg/common"
	"slg/control/account"
	"slg/control/army"
	"slg/control/capital"
	"slg/control/pve"
	"slg/control/unlock"
	"time"
)

type Route struct {
}

func (this *Route) Run(conn net.Conn) {
	// 初始化用户基础数据结构
	myData := &common.UserData{}
	myData.RefreshTime = time.Now().Unix()
	myData.Conn = &conn
	myData.Uid = 0

	// 设置定时器
	timeChan := make(chan int)
	defer close(timeChan)

	defer destructUser(myData, &timeChan)

	// 粘包处理缓存区域
	cacheData := &common.NetWorkCache{}
	cacheData.Len = 0
	cacheData.Cache = make([]byte, 10240)

	// 请求的接口标识
	var flag int = -1
	var id int = 0

	// 定义解析JSON的变量
	var parseData = make(map[string]interface{})
	var err error

	var outputData []byte = make([]byte, 1024)

	accountObj := &account.Account{}
	capitalObj := &capital.Capital{}
	unlockObj := &unlock.Unlock{}
	armyObj := &army.Army{}
	pveObj := &pve.Pve{}
	errorCount := 0

	go func(myData *common.UserData) {
		ticker := time.NewTicker(60 * time.Second)
		stop := 0
		for {
			select {
			case <-ticker.C:
				if myData.Uid > 0 {
					fmt.Println("开始定时器---校验用户数据", myData.Uid)
					capitalObj.AutoCollect(myData.Uid)
				} else if myData == nil {
					break
				}
			case <-timeChan:
				stop = 1
			}

			if stop == 1 {
				fmt.Println("消毁用户在线记录--跳出循环")
				break
			}
		}
	}(myData)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("---------------------------------------")
			fmt.Println(err)
			fmt.Println("---------------------------------------")
		}
	}()

	for {
		fmt.Println("ready")
		parseData, err = common.Request(conn, cacheData)
		fmt.Println("module", parseData["id"])

		if err != nil {

			if err.Error() == "EOF" || errorCount > 10 {
				fmt.Println("close error:", err.Error())
				conn.Close()
				goto Exit
			}
			errorCount++
			continue
		}

		if parseData == nil {
			continue
		}

		if pv, ok := parseData["id"].(uint64); ok {
			flag = int(math.Floor(float64(pv / 100)))
			id = int(pv)
		} else if pv, ok := parseData["id"].(float64); ok {
			flag = int(math.Floor(float64(pv / 100)))
			id = int(pv)
		} else {
			fmt.Println("id error:", parseData["id"])
			conn.Close()
			goto Exit
		}
		fmt.Println("flag", flag)

		switch flag {
		case CODE_ROUTE_PING:
			// 刷新心跳时间
			outputData = formatPong(id)

		case CODE_ROUTE_ACCOUNT:
			// 用户模块
			outputData = accountObj.Init(id, parseData, myData)

		case CODE_ROUTE_CAPITAL:
			// 主城模块
			outputData = capitalObj.Init(id, parseData, myData)

		case CODE_ROUTE_UNLOCK:
			outputData = unlockObj.Init(id, parseData, myData)

		case CODE_ROUTE_ARMY:
			outputData = armyObj.Init(id, parseData, myData)

		case CODE_ROUTE_PVE:
			outputData = pveObj.Init(id, parseData, myData)

		case 999:
			// 退出游戏
			conn.Close()
			goto Exit
		default:
			fmt.Println("default error")

			common.Response(conn, common.GetErrFormat(common.CODE_ERR_UNKNOW))
			conn.Close()
			goto Exit
		}

		if err != nil {
			fmt.Println("error", err)
			conn.Close()
			goto Exit
		}
		errorCount = 0
		fmt.Println("output:", outputData)
		common.Response(conn, outputData)
	}
Exit:
}

func destructUser(myData *common.UserData, timeChan *chan int) {
	if myData.Uid > 0 {
		// 清除用户集数据
		delete(common.OnlineUser, myData.Uid)
	}

	*timeChan <- 1

	fmt.Println("消毁用户在线记录")
}
