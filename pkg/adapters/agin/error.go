package agin

import (
	"nonoDemo/pkg/framework"
)

type HTTPError struct {
	framework.AppError
	code int
}

func NewHTTPError(err error, responseCode int) *HTTPError {
	if appErr, ok := err.(framework.AppError); ok {
		return &HTTPError{
			appErr,
			responseCode,
		}
	} else {
		return &HTTPError{
			framework.CommonInternalErr.Wrap(err, "internal error").(framework.AppError),
			responseCode,
		}
	}
}
