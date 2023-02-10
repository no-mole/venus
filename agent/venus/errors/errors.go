package errors

import (
	"errors"
)

var (
	ErrorLeaseExist                     = errors.New("grant lease exist")
	ErrorLeaseNotExist                  = errors.New("lease not exist")
	ErrorLeaseExpired                   = errors.New("lease expired")
	ErrorUserNotExist                   = errors.New("user not exit")
	ErrorUserNotExistOrPasswordNotMatch = errors.New("user not exit or password not match")
)
