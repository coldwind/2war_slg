package pve

import (
	"slg/common"
)

type Pve struct {
}

func (this *Pve) Init(id int, param interface{}, extra interface{}) (output []byte) {

	output = make([]byte, 1024)
	var errno int = 0

	switch id {
	case CODE_GET:
		output, errno = this.get(param, extra)
		break
	case CODE_BUILD:
		output, errno = this.build(param, extra)
		break
	case CODE_DISMANTLE:
		output, errno = this.dismantle(param, extra)
		break
	case CODE_ATTACK:
		output, errno = this.attack(param, extra)
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
