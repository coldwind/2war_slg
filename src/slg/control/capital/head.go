package capital

// 已有的产业
type FactoryStruct struct {
	Id          int `msgpack:"id"`
	Level       int `msgpack:"level"`
	CollectTime int `msgpack:"collectTime"`
}

type getStruct struct {
	Id   int              `msgpack:"id"`
	List []*FactoryStruct `msgpack:"list"`
	Time int64            `msgpack:"time"`
}

type collectStruct struct {
	Id          int `msgpack:"id"`
	Gold        int `msgpack:"gold"`
	Fid         int `msgpack:"fid"`
	CollectTime int `msgpack:"collectTime"`
	Time        int `msgpack:"time"`
}

type buyStruct struct {
	Id          int `msgpack:"id"`
	Gold        int `msgpack:"gold"`
	Fid         int `msgpack:"fid"`
	CollectTime int `msgpack:"collectTime"`
	Time        int `msgpack:"time"`
	Level       int `msgpack:"level"`
}
