package army

type armyStruct struct {
	Id  int `msgpack:"id"`
	Aid int `msgpack:"aid"`
	Num int `msgpack:"num"`
}

type getArmyStruct struct {
	Id      int                 `msgpack:"id"`
	CorpsId int                 `msgpack:"corpsId"`
	List    map[int]*armyStruct `msgpack:"list"`
	Time    int64               `msgpack:"time"`
}

type getCorpsStruct struct {
	Id   int   `msgpack:"id"`
	List []int `msgpack:"list"`
	Time int64 `msgpack:"time"`
}

type makeStruct struct {
	Id   int   `msgpack:"id"`
	Gold int   `msgpack:"gold"`
	Oil  int   `msgpack:"oil"`
	Time int64 `msgpack:"time"`
}
