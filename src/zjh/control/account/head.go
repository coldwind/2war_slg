package account

type LoginMPStruct struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Gold     int    `json:"gold"`
	Hp       int    `json:"hp"`
	Level    int    `json:"level"`
	Exp      int    `json:"exp"`
	Nickname string `json:"nickname"`
	Time     int64  `json:"time"`
}

type RegMPStruct struct {
	Id   int   `json:"id"`
	Time int64 `json:"time"`
}
