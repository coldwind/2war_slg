package hall

import (
	"zjh/common"
)

type Hall struct {
}

func (this *Hall) Init(id int, param interface{}, extra interface{}) (output []byte) {

	output = make([]byte, 1024)
	var errno int = 0

	switch id {
	case CODE_GET:
		output, errno = this.get(param, extra)
		break
	case CODE_CREATE:
		output, errno = this.create(param, extra)
		break
	case CODE_JOIN:
		output, errno = this.join(param, extra)
		break

	default:
		output = []byte{}
		break
	}

	if errno > 0 {
		return common.GetErrFormat(errno)
	}

	return
}
