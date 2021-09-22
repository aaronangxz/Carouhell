package models

const (
	CONSTANT_SUCCESS         = 0 //query ok
	CONSTANT_ERROR_PARAMS    = 1 //query parameters not valid
	CONSTANT_ERROR_NOT_FOUND = 2 //no result found
	CONSTANT_ERROR_DUPLICATE = 3 //for POST/PATCH/PUT apis, already exists
	CONSTANT_ERROR_OVERSIZE  = 4 //response exceeds max size
	CONSTANT_ERROR_TIMEOUT   = 5 //server timeout
	CONSTANT_ERROR_UNKNOWN   = 6 //etc
	CONSTANT_ERROR_DATABASE  = 7 //database error
)

func NewErrorParamdResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Parameter not match",
		ErrorCode: CONSTANT_ERROR_PARAMS,
	}
}

func NewNotFoundResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "No results",
		ErrorCode: CONSTANT_ERROR_NOT_FOUND,
	}
}

func NewErrorResponse(err error) ResponseMeta {
	return ResponseMeta{
		DebugMsg:  err.Error(),
		ErrorCode: CONSTANT_ERROR_DATABASE,
	}
}

func NewSuccessResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "",
		ErrorCode: CONSTANT_SUCCESS,
	}
}
