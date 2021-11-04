package contract

import "github.com/EricOgie/ope-be/ericerrors"

type BaseRequest interface {
	ValidateRequest() *ericerrors.EricError
}
