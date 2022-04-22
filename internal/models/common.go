package models

import "errors"

type UserIDParameter struct {
	User string `uri:"user" binding:"required"`
}

type OwnerParameter struct {
	TokenId string `uri:"tokenid" binding:"required"`
}

type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CustomErrorCode int32

const (
	InvalidUserIdParam CustomErrorCode = iota
	InvalidTokenID
	FailedToProcessRequest
)

func (e CustomErrorCode) String() string {
	switch e {
	case InvalidUserIdParam:
		return "INVALID_USER_ID_PARAMETER"
	case InvalidTokenID:
		return "INVALID_TOKEN_ID_PARAMETER"
	case FailedToProcessRequest:
		return "FAILED_TO_PROCESS_REQUEST"
	default:
		return "FAILED_TO_PROCESS_REQUEST"
	}
}

var (
	ErrFailedFetchData = errors.New("failed to fetch data")
)

type CheckHealthOK struct {
	Payload string `json:"body,omitempty"`
}
