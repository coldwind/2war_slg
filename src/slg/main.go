package main

import (
	"fmt"
	"slg/common"
	"slg/network"
)

func main() {

	// 创建运行时日志文件
	common.RunLogHandle = &common.LogSys{}
	common.RunLogHandle.InitLogFile("log/runtime.log")
	common.RunLogHandle.SaveLog("service started")
	common.OnlineUser = make(map[int]*common.UserData)
	confObj := &common.Conf{}

	// 读取游戏配置
	confObj.Init()

	// 初始化db
	common.DbHandle = &common.Dao{}
	common.DbHandle.InitConf()

	// 初始化redis
	common.RedisHandle = &common.RedisDao{}
	common.RedisHandle.InitConf()

	listen := network.Listen()

	defer listen.Close()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("---------------------------------------")
			fmt.Println(err)
			fmt.Println("---------------------------------------")
		}
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		route := &network.Route{}

		go route.Run(conn)
	}
}
