package function

import "github.com/pkg/errors"

var ErrDivByZero = errors.New("division by zero")

func Div(x, y int) (int,error)  {

	if y == 0 {
		return 0, ErrDivByZero
	}

	return x / y , nil
}
