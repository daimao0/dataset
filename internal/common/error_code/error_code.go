package error_code

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ErrorCode system error code
type ErrorCode struct {
	Code int
	Msg  string
	Err  error
}

var (
	FieldTypeNotExist = &ErrorCode{10002, "字段不存在", errors.New("字段不存在")}
)

func (errCode *ErrorCode) Error() string {
	bytes, err := json.MarshalIndent(errCode, "", "  ") // 使用两个空格作为缩进
	if err != nil {
		// 处理json.MarshalIndent可能发生的错误
		return fmt.Sprintf("{\"error\": \"%v\"}", err)
	}
	return string(bytes)
}
