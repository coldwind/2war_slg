package unlock

import (
	"slg/common"
)

type Unlock struct {
}

func (this *Unlock) Init(id int, param interface{}, extra interface{}) (output []byte) {

	output = make([]byte, 1024)
	var errno int = 0

	switch id {
	case CODE_GET:
		output, errno = this.get(param, extra)
		break

	case CODE_UNLOCK:
		output, errno = this.unlock(param, extra)
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
