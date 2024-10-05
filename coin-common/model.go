// Package common
// @Author twilikiss 2024/4/28 22:21:21
package common

type ElysiaCode int

const SuccessCode ElysiaCode = 0
const DefaultFailCode ElysiaCode = -999

type Result struct {
	Code    ElysiaCode `json:"code"`
	Message string     `json:"message"`
	Data    any        `json:"data"`
}

//type error struct {
//	*error
//	Code int
//}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Fail(code ElysiaCode, msg string) {
	r.Code = code
	r.Message = msg
}

func (r *Result) Success(data any) {
	r.Code = SuccessCode
	r.Message = "success"
	r.Data = data
}

func (r *Result) Deal(data any, err error, errorCode ElysiaCode) *Result {
	if err != nil {
		if errorCode == -1 {
			errorCode = DefaultFailCode
		}
		r.Fail(errorCode, err.Error())
		return r
	}
	r.Success(data)
	return r
}
