package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
}

func (this *Conf) Init() {

	// PVE地块类型配置
	groundTypeBuf, err := ioutil.ReadFile("conf/groundType.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	GroundTypeConfs = make(map[int]*GroundTypeStruct)
	groundTypeTemp := make([]*GroundTypeStruct, 0, 1024)
	err = json.Unmarshal(groundTypeBuf, &groundTypeTemp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, v := range groundTypeTemp {
		GroundTypeConfs[v.Id] = v
	}

	// PVE地图配置
	pveMapBuf, err := ioutil.ReadFile("conf/pveMap.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	PveMapConfs = make(map[int]*PveMapStruct)
	pveMapTemp := make([]*PveMapStruct, 0, 1024)
	err = json.Unmarshal(pveMapBuf, &pveMapTemp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, v := range pveMapTemp {
		PveMapConfs[v.Id] = v
	}

	// 工厂配置
	produceBuf, err := ioutil.ReadFile("conf/produce.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ProduceConfs = make(map[int]*ProduceStruct)
	err = json.Unmarshal(produceBuf, &ProduceConfs)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 用户注册初始化数据配置
	initUserBuf, err := ioutil.ReadFile("conf/initUser.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	InitUserConf = &InitUserStruct{}
	err = json.Unmarshal(initUserBuf, InitUserConf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// INIT PVEMAP
	InitPveMapConf = make([]int, 0, 10)
	initPveMapBuf, err := ioutil.ReadFile("conf/initPveMap.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(initPveMapBuf, &InitPveMapConf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 解锁经理人
	unlockManageBuf, err := ioutil.ReadFile("conf/unlockManage.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	UnlockManageConfs = make(map[int]*UnlockManageStruct)
	unlockManageTemp := make([]*UnlockManageStruct, 0, 1024)
	err = json.Unmarshal(unlockManageBuf, &unlockManageTemp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, v := range unlockManageTemp {
		UnlockManageConfs[v.Id] = v
	}

	// 解锁部队
	unlockCorpsBuf, err := ioutil.ReadFile("conf/corps.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	UnlockCorpsConfs = make(map[int]*UnlockCorpsStruct)
	unlockCorpsTemp := make([]*UnlockCorpsStruct, 0, 10)
	err = json.Unmarshal(unlockCorpsBuf, &unlockCorpsTemp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, v := range unlockCorpsTemp {
		UnlockCorpsConfs[v.Id] = v
	}

	// 军队配置
	ArmyBuf, err := ioutil.ReadFile("conf/army.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	armyTemp := make([]*ArmyStruct, 0, 1024)

	err = json.Unmarshal(ArmyBuf, &armyTemp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	ArmyConfs = make(map[int]*ArmyStruct)
	for _, v := range armyTemp {
		ArmyConfs[v.Id] = v
	}

	// NPC部队
	NpcCorpsConfs = make([]*NpcCorps, 0, 1024)
	NpcCorpsBuf, err := ioutil.ReadFile("conf/npcCorps.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(NpcCorpsBuf, &NpcCorpsConfs)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
