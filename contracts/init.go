package contracts

// BaseRequest defines basic Request structure for all APIs.
type BaseRequest struct {
	Method    *string `json:"-"`
	RequestID *string `json:"-"`
}

// BaseResponse defines basic Response for all APIs
type BaseResponse struct {
	RequestID *string    `json:"request_id"`
	Method    *string    `json:"method"`
	HTTPCode  *int       `json:"http_code"`
	ErrorData *ErrorData `json:"error_data"`
}

// ErrorData defines the Error structure for response
type ErrorData struct {
	Code        uint64 `json:"code"`
	Description string `json:"description"`
}

func Init() error {
	return nil
}
