package forth

import "errors"

const (
	NO_ERROR = iota
	ERR_STACK_UNDERFLOW
	ERR_DIVISION_BY_ZERO
	ERR_REDEF_NUMBER
	ERR_UNDEF_WORD
)

func getError() error {
	var err error
	switch errorCode {
	case NO_ERROR:
		err = nil
	case ERR_STACK_UNDERFLOW:
		err = errors.New("stack underflow")
	case ERR_DIVISION_BY_ZERO:
		err = errors.New("division by zero")
	case ERR_REDEF_NUMBER:
		err = errors.New("redefining number")
	case ERR_UNDEF_WORD:
		err = errors.New("undefined word")
	}

	return err
}
