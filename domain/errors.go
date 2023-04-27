package domain

type ErrorCode string

const (
	ErrInvalidInput ErrorCode = "INVALID_INPUT" // 400
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"  // 401
	ErrForbidden    ErrorCode = "FORBIDDEN"     // 403
	ErrNotFound     ErrorCode = "NOT_FOUND"     // 404
	ErrInternal     ErrorCode = "INTERNAL"      // 500
)
