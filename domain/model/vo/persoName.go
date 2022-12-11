package vo

import (
	"fmt"
)

type PersonName string

func NewName(value string) (*PersonName, error) {
	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name newName()")
		return nil, err
	}

	name := PersonName(value)

	return &name, nil
}
