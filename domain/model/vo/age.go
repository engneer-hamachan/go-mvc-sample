package vo

import (
	"fmt"
)

type Age int

func NewAge(value int) (*Age, error) {
	if value >= 100 {
		err := fmt.Errorf("%s", "invalid age arg:age NewAge()")
		return nil, err
	}

	age := Age(value)

	return &age, nil
}
