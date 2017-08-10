package army

import (
	"slg/common"
)

type Army struct {
}

func (this *Army) Init(id int, param interface{}, extra interface{}) (output []byte) {

	output = make([]byte, 1024)
	var errno int = 0

	switch id {
	case CODE_GET_ALL:
		//output, errno = this.getAll(param, extra)
		break

	case CODE_GET_CORPS:
		output, errno = this.getCorps(param, extra)
		break

	case CODE_GET_ARMY_CORPS:
		output, errno = this.getByCorps(param, extra)
		break

	case CODE_MAKE_ARMY:
		output, errno = this.makeArmy(param, extra)
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
