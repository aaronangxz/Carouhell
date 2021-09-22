package models

const (
	CONSTANT_SUCCESS         = 0 //query ok
	CONSTANT_ERROR_PARAMS    = 1 //query parameters not valid
	CONSTANT_ERROR_JSON      = 2 //query parameters does not match JSON
	CONSTANT_ERROR_NOT_FOUND = 3 //no result found
	CONSTANT_ERROR_DUPLICATE = 4 //for POST/PATCH/PUT apis, already exists
	CONSTANT_ERROR_OVERSIZE  = 5 //response exceeds max size
	CONSTANT_ERROR_TIMEOUT   = 6 //server timeout
	CONSTANT_ERROR_UNKNOWN   = 7 //etc
	CONSTANT_ERROR_DATABASE  = 8 //database error
)

func NewSuccessResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "",
		ErrorCode: CONSTANT_SUCCESS,
	}
}

func NewParamErrorsResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Parameter not match",
		ErrorCode: CONSTANT_ERROR_PARAMS,
	}
}

func NewJSONErrorResponse(err error) ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "JSON Error: " + err.Error(),
		ErrorCode: CONSTANT_ERROR_JSON,
	}
}

func NewNotFoundResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "No results.",
		ErrorCode: CONSTANT_ERROR_NOT_FOUND,
	}
}

func NewDuplicateErrorResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Record already exists.",
		ErrorCode: CONSTANT_ERROR_DUPLICATE,
	}
}

func NewOversizeErrorResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Result is too large to return.",
		ErrorCode: CONSTANT_ERROR_OVERSIZE,
	}
}

func NewTimeoutErrorResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Timeout Error.",
		ErrorCode: CONSTANT_ERROR_TIMEOUT,
	}
}

func NewUnknownErrorResponse() ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Unknown Error.",
		ErrorCode: CONSTANT_ERROR_UNKNOWN,
	}
}

func NewDBErrorResponse(err error) ResponseMeta {
	return ResponseMeta{
		DebugMsg:  "Database Error: " + err.Error(),
		ErrorCode: CONSTANT_ERROR_DATABASE,
	}
}
