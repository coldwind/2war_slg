package unlock

type getStruct struct {
	Id   int   `msgpack:"id"`
	Type int   `msgpack:"type"`
	List []int `msgpack:"list"`
	Time int64 `msgpack:"time"`
}

type unlockStruct struct {
	Id   int   `msgpack:"id"`
	Type int   `msgpack:"type"`
	Gold int   `msgpack:"gold"`
	Mid  int   `msgpack:"mid"`
	Time int64 `msgpack:"time"`
}
