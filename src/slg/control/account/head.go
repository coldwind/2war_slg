package account

type LoginMPStruct struct {
	Id       int    `msgpack:"id"`
	Uid      int    `msgpack:"uid"`
	Gold     int    `msgpack:"gold"`
	Hp       int    `msgpack:"hp"`
	Level    int    `msgpack:"level"`
	Exp      int    `msgpack:"exp"`
	MainCity int    `msgpack:"mainCity"`
	Nickname string `msgpack:"nickname"`
	Time     int64  `msgpack:"time"`
	City     []int  `msgpack:"city"` // 其他城市ID集合

}

type RegMPStruct struct {
	Id   int   `msgpack:"id"`
	Time int64 `msgpack:"time"`
}
