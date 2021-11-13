package errori18n

type ErrorCode string

// Auto-generated using `go run errori18n/codesgen.go`. Please do not edit this file.
const (
	ERR_UNKNOWN        ErrorCode = "ERR_UNKNOWN"
	ERR_SIMPLE         ErrorCode = "ERR_SIMPLE"
	ERR_VARIADIC       ErrorCode = "ERR_VARIADIC"
	ERR_DIVIDE_BY_ZERO ErrorCode = "ERR_DIVIDE_BY_ZERO"
)