package network

import "slg/common"

type PongMPStruct struct {
	Id int `msgpack:"id"`
}

func formatPong(id int) (output []byte) {
	pong := &PongMPStruct{}
	pong.Id = 0

	output, err := common.Serialize(pong)
	if err != nil {
		return []byte{}
	}

	return output
}
