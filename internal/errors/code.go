package errors

type Code string

const (
	CodeValidation   Code = "VALIDATION_ERROR"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeForbidden    Code = "FORBIDDEN"
	CodeNotFound     Code = "NOT_FOUND"
	CodeConflict     Code = "CONFLICT"
	// Server errors
	CodeInternal Code = "INTERNAL_ERROR"
	CodeTimeout  Code = "TIMEOUT"
	CodeExternal Code = "EXTERNAL_SERVICE_ERROR"
)
