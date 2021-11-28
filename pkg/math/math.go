package math

import "errors"

type IMath interface {
	Plus(a int, b int) int
	Minus(a int, b int) int
	Multiply(a int, b int) int
	Divide(a int, b int) (result int, err error)
}

type Math struct {
	math *IMath
}

func (m *Math) Plus(a int, b int) int {
	return a + b
}

func (m *Math) Minus(a int, b int) int {
	return a - b
}

func (m *Math) Multiply(a int, b int) int {
	return a * b
}

func (m *Math) Divide(a int, b int) (result int, err error) {

	if b == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return a / b, nil
}
