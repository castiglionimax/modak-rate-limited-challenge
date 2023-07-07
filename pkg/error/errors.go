package error

import "errors"

var ErrRateLimitedReach = errors.New("rate limited has reached")

type (
	HandlerError struct {
		Cause error
	}

	NotFoundError struct {
		Cause error
	}
)

func (e HandlerError) Error() string {
	return e.Cause.Error()
}

func (e NotFoundError) Error() string {
	return e.Cause.Error()
}
