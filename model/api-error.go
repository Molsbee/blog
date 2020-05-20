package model

type ApiError interface {
	Error() string
	StatusCode() int
	Details() map[string]interface{}
}

type apiError struct {
	S int                    `json:"-"`
	M string                 `json:"error,omitempty"`
	D map[string]interface{} `json:"details,omitempty"`
}

type apiErrorBuilder struct {
	statusCode int
	message    string
	details    map[string]interface{}
}

func ErrorBuilder() apiErrorBuilder {
	return apiErrorBuilder{}
}

func (a apiErrorBuilder) StatusCode(statusCode int) apiErrorBuilder {
	a.statusCode = statusCode
	return a
}

func (a apiErrorBuilder) Message(message string) apiErrorBuilder {
	a.message = message
	return a
}

func (a apiErrorBuilder) Details(details map[string]interface{}) apiErrorBuilder {
	a.details = details
	return a
}

func (a apiErrorBuilder) Build() ApiError {
	return apiError{
		S: a.statusCode,
		M: a.message,
		D: a.details,
	}
}

func (a apiError) Error() string {
	return a.M
}

func (a apiError) StatusCode() int {
	return a.S
}

func (a apiError) Details() map[string]interface{} {
	return a.D
}
