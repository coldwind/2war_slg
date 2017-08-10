package common

import (
	"github.com/gorilla/websocket"
)

// 用户数据结构
type UserData struct {
	Uid         int             `json:"uid"`
	Conn        *websocket.Conn `json:",omitempty"`
	RefreshTime int64           `json:"refreshTime"`
	Nickname    string          `json:"nickname"`
	Hp          int             `json:"hp"`
	Gold        int             `json:"gold"`
	Level       int             `json:"level"`
	Exp         int             `json:"exp"`
}

// 大厅数据结构
type HallInfo struct {
	HallId int         `json:"hallId"`
	Name   string      `json:"name"`
	Uid    int         `json:"uid"`
	Users  []*UserData `json:"users"`
}
