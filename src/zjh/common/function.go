package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

func GetErrFormat(id int) (output []byte) {
	format := &ErrorFormat{}
	format.Id = id

	output, err := Serialize(format)
	if err != nil {
		return []byte{}
	}

	return output
}

func Serialize(Data interface{}) (serialize []byte, err error) {
	serialize, err = json.Marshal(Data)

	if err != nil {
		return []byte{}, errors.New("serialize error")
	}

	return serialize, err
}

func Md5Encode(data string) (output string) {
	h := md5.New()
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

func Response(conn *websocket.Conn, outputData []byte) bool {

	err := conn.WriteMessage(websocket.TextMessage, outputData)
	if err != nil {
		log.Println("write:", err)
		return false
	}

	return true
}

func SetBitFlag(data uint, bit uint) uint {
	if bit < 1 {
		bit = 1
	}

	data |= 1 << (bit - 1)

	return data
}

func GetBitFlag(data uint, bit uint) uint {
	data &= 1 << (bit - 1)

	return data
}

func CheckLogin(loginData *UserData) bool {
	if loginData.Uid > 0 {
		return true
	}

	return false
}
