package ssdb

import (
	"fmt"
)

var errIsEmpty = fmt.Errorf("error: respone is empty")

func ResultProcessing(args []interface{}, resp Values, err error) (Values, error) {

	if err != nil {
		return nil, fmt.Errorf("%s: %v", err.Error(), args)
	}

	if len(resp) < 1 {
		return nil, errIsEmpty
	}

	if resp[0].Equal(ok) {
		return resp[1:], nil
	}

	if resp[0].Equal(notFound) {
		return nil, nil
	}

	if resp[0].Equal(clientError) {
		return resp[1:], fmt.Errorf("client error: %v", resp)
	}

	if len(args) > 0 {
		return nil, fmt.Errorf("error: %v, parameter %v", resp[0], args)
	}
	return nil, fmt.Errorf("error: parameter %v", resp[0])
}
