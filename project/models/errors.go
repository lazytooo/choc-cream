package models

import "errors"

const ErrCodeOK = 1000

var (
	ErrInvalidParam = errors.New("参数错误")
	ErrDB           = errors.New("服务繁忙，请稍后重试")
)

// nolint: gocyclo
func GetErrorCode(err error) int32 {
	switch err {
	case nil:
		return ErrCodeOK
	case ErrInvalidParam:
		return ErrCodeOK + 1
	case ErrDB:
		return ErrCodeOK + 2
	default:
		return -1
	}
}

func GetErrorMap(err error) map[string]interface{} {
	var msg = "OK"
	if err != nil {
		msg = err.Error()
	}

	return map[string]interface{}{
		"code": GetErrorCode(err),
		"msg":  msg,
	}
}
