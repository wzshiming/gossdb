package ssdb

import (
	"fmt"
)

var (
	ErrIsEmpty  = fmt.Errorf("error: respone is empty")
	ErrNotFound = fmt.Errorf("error: not found")
)

func (c Client) ResultProcessing(resp Values, err error) (Values, error) {

	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	if len(resp) < 1 {
		return nil, ErrIsEmpty
	}

	if resp[0].Equal(ok) && len(resp) >= 2 {
		return resp[1:], nil
	}

	if resp[0].Equal(notFound) {
		if c.ignoreGetNotFoundError {
			return nil, nil
		}
		return nil, ErrNotFound
	}

	if resp[0].Equal(clientError) && len(resp) >= 2 {
		return nil, fmt.Errorf("error: client error: %v", resp[1:].String())
	}

	return nil, fmt.Errorf("error: parameter: %v", resp[0])
}
