package errors

const (
	EmailAlreadyUsed    = "EMAIL_ALREADY_USED"
	UsernameAlreadyUsed = "USERNAME_ALREADY_USED"
	NotFound            = "NOT_FOUND"
	BadQuery            = "BAD_QUERY"
	Forbidden           = "FORBIDDEN"
	Unauthorized        = "UNAUTHORIZED"
	WrongCredentials    = "WRONG_CREDENTIALS"
	ServerError         = "INTERNAL_SERVER_ERROR"
	Validation          = "JSON_VALIDATE_ERROR"
)

type HttpErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

type HttpError struct {
	StatusCode int         `json:"status_code,omitempty"`
	Key        string      `json:"key,omitempty"`
	Stack      interface{} `json:"stack"`
}

func (err HttpError) Status() int {
	return err.StatusCode
}

func (err HttpError) Error() string {
	return err.Key
}

func (err HttpError) Causes() interface{} {
	return err.Stack
}

func NewHttpError(status int, key string, stack interface{}) *HttpError {
	return &HttpError{
		StatusCode: status,
		Key:        key,
		Stack:      stack,
	}
}
