package defs

type Err struct {
	Error string
	ErrorCode string
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC:400,Error:Err{Error:"Bad request body",ErrorCode:"001"}}

	ErrorNotAuthUser = ErrResponse{HttpSC:403,Error:Err{Error:"User authentication failed",ErrorCode:"002"}}
	ErrorDBError = ErrResponse{HttpSC:500,Error:Err{Error:"DB failed",ErrorCode:"003"}}
	ErrorInternalFaults = ErrResponse{HttpSC:500,Error:Err{Error:"Internal service error",ErrorCode:"004"}}
)
