package capital

import (
	"slg/common"
)

type Capital struct {
}

func (this *Capital) Init(id int, param interface{}, extra interface{}) (output []byte) {

	common.DbHandle.TableName = "user"

	output = make([]byte, 1024)
	var errno int = 0

	switch id {
	case CODE_GET:
		output, errno = this.get(param, extra)
		break

	case CODE_COLLECT:
		output, errno = this.collect(param, extra)
		break

	case CODE_BUY_FACTORY:
		output, errno = this.buyFactory(param, extra)
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
