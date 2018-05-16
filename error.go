package ssdb

import (
	"fmt"
)

func makeError(err error, resp Values, args []interface{}) error {
	if err != nil {
		fmt.Errorf("%s: %v", err.Error(), args)
	}
	if len(resp) < 1 {
		return fmt.Errorf("error: respone is empty.")
	}

	if resp[0].Equal(notFound) {
		return nil
	}

	if len(args) > 0 {
		return fmt.Errorf("error: %v, parameter %v", resp[0], args)
	}
	return fmt.Errorf("error: parameter %v", resp[0])
}
