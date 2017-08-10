package hall

import (
	"zjh/common"
)

type HallListMP struct {
	Id   int                `json:"id"`
	List []*common.HallInfo `json:"list"`
	Time int64              `json:"time"`
}

type CreateHallMP struct {
	Id     int   `json:"id"`
	HallId int   `json:"hallId"`
	Time   int64 `json:"time"`
}

type JoinMP struct {
	Id       int              `json:"id"`
	HallInfo *common.HallInfo `json:"hallInfo"`
	Time     int64            `json:"time"`
}
