package common

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"log"
	"net"

	"gopkg.in/vmihailenco/msgpack.v2"
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
	serialize, err = msgpack.Marshal(Data)

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

func Response(conn net.Conn, outputData []byte) bool {
	var outputLen int16
	outputLen = int16(len(outputData))
	resByte := make([]byte, 0)

	resBuf := bytes.NewBuffer([]byte{})
	binary.Write(resBuf, binary.BigEndian, outputLen)
	resByte = append(resByte, []byte("xtrops")...)

	if outputLen > 0 {
		resByte = append(append(resByte, resBuf.Bytes()...), outputData...)
		log.Println(string(resByte), "---end---")
	}

	n, err := conn.Write(resByte)
	log.Println("write:", err, n)

	return true
}

func Request(conn net.Conn, cacheData *NetWorkCache) (mpData map[string]interface{}, err error) {
	requestData := make([]byte, 1024)
	log.Println("requested", cacheData.Len)

	if cacheData.Len > 0 {
		mpData = make(map[string]interface{})

		log.Println("source data:", cacheData.Cache)

		head := cacheData.Cache[0:6]
		log.Println("head:", string(head))

		if string(head) != "xtrops" {
			return mpData, errors.New("head error")
		}

		if len(cacheData.Cache) <= 8 {
			return mpData, errors.New("too short")
		}

		byteLen := int(binary.BigEndian.Uint16(cacheData.Cache[6:8]))
		log.Println("byteLen:", byteLen, len(cacheData.Cache))

		if len(cacheData.Cache) < byteLen+8 {
			return mpData, errors.New("too short")
		}
		log.Println("mpdata:", cacheData.Cache[8:byteLen+8])

		err = msgpack.Unmarshal(cacheData.Cache[8:byteLen+8], &mpData)
		if err != nil {
			return mpData, err
		}

		cacheData.Cache = cacheData.Cache[8+byteLen:]
		cacheData.Len = cacheData.Len - 8 - int(byteLen)
		log.Println("request data:", mpData)

		return mpData, nil

	}
	log.Println("request data")

	n, err := conn.Read(requestData)
	if err != nil {
		return mpData, err
	}

	requestData = requestData[:n]
	log.Println("requested data", n)

	// 计算头内容
	if uint16(n) > TRANS_MAX_LEN {
		return mpData, errors.New("too long")
	}

	cacheData.Cache = append(cacheData.Cache[0:cacheData.Len], requestData...)
	cacheData.Len += len(requestData)

	return nil, nil
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
