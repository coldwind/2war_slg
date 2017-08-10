package account

import (
	"zjh/common"
)

type Account struct {
}

func (this *Account) Init(id int, param interface{}, extra interface{}) (output []byte) {

	common.DbHandle.TableName = "user"

	output = make([]byte, 1024)
	var errno int = 0

	switch id {
	case CODE_LOGIN:
		output, errno = this.login(param, extra)
		break

	case CODE_REG:
		output, errno = this.reg(param, extra)
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
