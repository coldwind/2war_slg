package pve

type Pos struct {
	X int `msgpack:"x"`
	Y int `msgpack:"y"`
}

type CorpsInfo struct {
	Spos int `msgpack:"spos"`
	Epos int `msgpack:"epos"`
}

type CityInfo struct {
	Id  int `msgpack:"id"`
	Cid int `msgpack:"cid"`
	X   int `msgpack:"x"`
	Y   int `msgpack:"y"`
}

type GetStruct struct {
	Id   int               `msgpack:"id"`
	City map[int]*CityInfo `msgpack:"city"`
	MI   []int             `msgpack:"mi"`
	CI   []CorpsInfo       `msgpack:"ci"`
	Time int64             `msgpack:"time"`
}

type BuildStruct struct {
	Id     int   `msgpack:"id"`
	CityId int   `msgpack:"cityId"`
	Hp     int   `msgpack:"hp"`
	Time   int64 `msgpack:"time"`
}

type DismantleStruct struct {
	Id     int   `msgpack:"id"`
	CityId int   `msgpack:"cityId"`
	Time   int64 `msgpack:"time"`
}

type AttackInfoStruct struct {
	CityId  int   `json:"cityId"`
	CorpsId int   `json:"corpsId"`
	Spos    int   `json:"spos"`
	Epos    int   `json:"epos"`
	Atime   int64 `json:"atime"`
}

type AttackInfoMpStruct struct {
	Id          int               `msgpack:"id"`
	CityId      int               `msgpack:"cityId"`
	CorpsId     int               `msgpack:"corpsId"`
	Spos        int               `msgpack:"spos"`
	Epos        int               `msgpack:"epos"`
	Win         int               `msgpack:"win"`
	MyArmy      []*FightFormation `msgpack:"myArmy"`
	EnemyArmy   []*FightFormation `msgpack:"enemyArmy"`
	FightRecord []*FightRecord    `msgpack:"fightRecord"`
	Atime       int64             `msgpack:"atime"`
	Time        int64             `msgpack:"time"`
}

// PVE战斗记录
type FightRecord struct {
	Atk     int           `msgpack:"atk"`
	Targets []*HurtRecord `msgpack:"target"`
}

// 伤害记录
type HurtRecord struct {
	Target int `msgpack:"target"`
	Hurt   int `msgpack:"hurt"`
}

// 战斗阵型
type FightFormation struct {
	Aid int `msgpack:"aid"`
	Pos int `msgpack:"pos"`
	Num int `msgpack:"num"`
}
