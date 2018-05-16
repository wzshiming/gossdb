package ssdb

import (
	"fmt"
)

var errIsEmpty = fmt.Errorf("error: respone is empty")

func makeError(err error, resp Values, args []interface{}) error {
	if err != nil {
		return fmt.Errorf("%s: %v", err.Error(), args)
	}

	if len(resp) < 1 {
		return errIsEmpty
	}

	if resp[0].Equal(notFound) {
		return nil
	}

	if len(args) > 0 {
		return fmt.Errorf("error: %v, parameter %v", resp[0], args)
	}
	return fmt.Errorf("error: parameter %v", resp[0])
}
