package common

import (
	"net"
)

/**************** REDIS数据结构 ******************/
// 工厂数据结构
type Factory struct {
	Id          int
	Level       int
	CollectTime int
}

type GlobalUnlock struct {
	City map[int][]int
}

// 用户数据结构
type UserData struct {
	Uid         int       `msgpack:"uid"`
	Conn        *net.Conn `msgpack:",omitempty"`
	RefreshTime int64     `msgpack:"refreshTime"`
	Nickname    string    `msgpack:"nickname"`
	Hp          int       `msgpack:"hp"`
	Gold        int       `msgpack:"gold"`
	Oil         int       `msgpack:"oil"`
	Level       int       `msgpack:"level"`
	MainCity    int       `msgpack:"mainCity"`
	Exp         int       `msgpack:"exp"`
}

/***************** 配置文件结构 *****************/
// 地块类型
type GroundTypeStruct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Income      int    `json:"income"`
	Oil         int    `json:"oil"`
	PveHp       int    `json:"pveHp"`
	PvpHp       int    `json:"pvpHp"`
	PveCorpsId  int    `json:"pveCorpsId"`
	PveCorpsNum int    `json:"pveCorpsNum"`
	PvpCorpsId  int    `json:"pvpCorpsId"`
	PvpCorpsNum int    `json:"pvpCorpsNum"`
	NpcRecover  int    `json:"npcRecover"`
	Desc        string `json:"desc"`
}

// PVE地图
type PveMapStruct struct {
	Id   int `json:"id"`
	Type int `json:"type"`
	X    int `json:"x"`
	Y    int `json:"y"`
}

// 工厂
type ProduceStruct struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Price            int     `json:"price"`
	Para             int     `json:"para"`
	CostPara         float32 `json:"cost_para"`
	Interval         int     `json:"interval"`
	Para2            int     `json:"para2"`
	Produce          int     `json:"produce"`
	Para3            int     `json:"para3"`
	InitProductivity float32 `json:"init_productivity"`
	Para100          float32 `json:"para100"`
	LevelLimit       int     `json:"level_limit"`
}

// Init--User
type InitUserStruct struct {
	Gold     int `json:"gold"`
	Level    int `json:"level"`
	Exp      int `json:"exp"`
	Hp       int `json:"hp"`
	Oil      int `json:"oil"`
	MainCity int `json:"mainCity"`
	CityX    int `json:"cityX"`
	CityY    int `json:"cityY"`
}

// 解锁经理
type UnlockManageStruct struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	GoldAddition int    `json:"goldAddition"`
	UnlockPrice  int    `json:"unlockPrice"`
	FactoryId    int    `json:"factoryId"`
	PrevId       int    `json:"prevId"`
	NextId       int    `json:"nextId"`
}

// 解锁部队
type UnlockCorpsStruct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	UnlockPrice int    `json:"unlockPrice"`
	PrevId      int    `json:"prevId"`
	NextId      int    `json:"nextId"`
}

// 军队
type ArmyStruct struct {
	Id          int    `json:"id"`
	ArmyId      int    `json:"armyId"`
	Name        string `json:"name"`
	MakePrice   int    `json:"makePrice"`
	MakeOil     int    `json:"makeOil"`
	Level       int    `json:"level"`
	Atk         int    `json:"atk"`
	Def         int    `json:"def"`
	Spd         int    `json:"spd"`
	Pace        int    `json:"pace"`
	MinRange    int    `json:"minRange"`
	MaxRange    int    `json:"maxRange"`
	InstantFire int    `json:"instantFire"`
	AtkType     int    `json:"atkType"`
	DefType     int    `json:"defType"`
	UnlockPrice int    `json:"unlockPrice"`
	UnlockOil   int    `json:"unlockOil"`
	PrevId      int    `json:"prevId"`
	NextId      int    `json:"nextId"`
}

// NPC部队
type NpcCorps struct {
	CorpsId int `json:"corpsId"`
	Aid     int `json:"aid"`
	Num     int `json:"num"`
	Aid1    int `json:"aid1"`
	Num1    int `json:"num1"`
	Aid2    int `json:"aid2"`
	Num2    int `json:"num2"`
	Aid3    int `json:"aid3"`
	Num3    int `json:"num3"`
	Aid4    int `json:"aid4"`
	Num4    int `json:"num4"`
	Hero    int `json:"hero"`
}
