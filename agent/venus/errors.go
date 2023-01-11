package venus

import (
	"errors"
)

var (
	ErrorLeaseExist    = errors.New("grant lease exist")
	ErrorLeaseNotExist = errors.New("lease not exist")
	ErrorLeaseExpired  = errors.New("lease expired")
)
