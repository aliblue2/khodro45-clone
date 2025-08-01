package helper

import "github.com/aliblue2/khodro45/api/validations"

type BaseResponse struct {
	Result        any                            `json:"result"`
	Success       bool                           `json:"success"`
	ResultCode    int                            `json:"result-code"`
	Error         error                          `json:"error"`
	ValidationErr *[]validations.ValidationError `json:"validation-error"`
}

func BaseResponseHandler(result any, success bool, resultCode int) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}
func BaseResponseHandlerWithError(result any, success bool, resultCode int, err error) *BaseResponse {
	return &BaseResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}
func BaseResponseHandlerWithValidationError(result any, success bool, resultCode int, err error) *BaseResponse {
	return &BaseResponse{
		Result:        result,
		Success:       success,
		ResultCode:    resultCode,
		ValidationErr: validations.GetValidationError(err),
	}
}
