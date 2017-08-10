package main

import (
	"fmt"
	"zjh/common"
	"zjh/network"
)

func main() {

	// 创建运行时日志文件
	common.RunLogHandle = &common.LogSys{}
	common.RunLogHandle.InitLogFile("log/runtime.log")
	common.RunLogHandle.SaveLog("service started")
	common.OnlineUser = make(map[int]*common.UserData)
	common.HallData = make(map[int]*common.HallInfo)
	common.HallCount = 1

	// 初始化db
	common.DbHandle = &common.Dao{}
	common.DbHandle.InitConf()

	// 初始化redis
	common.RedisHandle = &common.RedisDao{}
	common.RedisHandle.InitConf()

	//defer listen.Close()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("---------------------------------------")
			fmt.Println(err)
			fmt.Println("---------------------------------------")
		}
	}()

	network.Listen()
}
