package common

type ErrorCode int

const (
	ErrUnknown ErrorCode = iota + 1
	ErrNotFound
	ErrInvalidInput
	ErrPermissionDenied
	ErrDatabaseError
)

func (e ErrorCode) String() string {
	switch e {
	case ErrNotFound:
		return "USER_NOT_FOUND"
	case ErrInvalidInput:
		return "INVALID_INPUT"
	case ErrPermissionDenied:
		return "PERMISSION_DENIED"
	case ErrDatabaseError:
		return "DATABASE_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}

func (e ErrorCode) HttpStatus() int {
	switch e {
	case ErrInvalidInput:
		return 400
	case ErrNotFound:
		return 404
	case ErrPermissionDenied:
		return 403
	case ErrDatabaseError:
		return 500
	default:
		return 500
	}
}

func (e ErrorCode) StatusCode() string {
	switch e {
	case ErrNotFound:
		return "01"
	case ErrInvalidInput:
		return "02"
	case ErrPermissionDenied:
		return "03"
	case ErrDatabaseError:
		return "04"
	default:
		return "99"
	}
}
