package tools

import "strconv"

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, codeMsg ...interface{}) {
	if !condition {
		statusCode := 200
		msg := ""
		if len(codeMsg) > 0 {
			statusCode = codeMsg[0].(int)
		}
		if len(codeMsg) > 1 {
			msg = codeMsg[1].(string)
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasError(err error, codeMsg ...interface{}) {
	if err != nil {
		statusCode := 500
		msg := err.Error()
		if len(codeMsg) > 0 {
			statusCode = codeMsg[0].(int)
		}

		if len(codeMsg) > 1 {
			msg = codeMsg[1].(string)
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}
