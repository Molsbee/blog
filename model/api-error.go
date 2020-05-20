package model

import "strings"

type ApplicationError interface {
	Error() string
	StatusCode() int
	Details() map[string]interface{}
}

type ApiError struct {
	S int                    `json:"-"`
	M string                 `json:"error,omitempty"`
	D map[string]interface{} `json:"details,omitempty"`
}

func (a ApiError) Error() string {
	return a.M
}

func (a ApiError) StatusCode() int {
	return a.S
}

func (a ApiError) Details() map[string]interface{} {
	return a.D
}

type apiErrorBuilder struct {
	statusCode int
	message    string
	details    map[string]interface{}
}

func ErrorBuilder() apiErrorBuilder {
	return apiErrorBuilder{
		details: make(map[string]interface{}),
	}
}

func (a apiErrorBuilder) StatusCode(statusCode int) apiErrorBuilder {
	a.statusCode = statusCode
	return a
}

func (a apiErrorBuilder) Message(message string) apiErrorBuilder {
	a.message = message
	return a.AddDetail("message", message)
}

func (a apiErrorBuilder) AddDetail(key string, value interface{}) apiErrorBuilder {
	a.details[strings.ToLower(key)] = value
	return a
}

func (a apiErrorBuilder) Build() ApplicationError {
	return ApiError{
		S: a.statusCode,
		M: a.message,
		D: a.details,
	}
}
