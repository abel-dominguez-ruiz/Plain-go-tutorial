package math

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type MathMock struct {
}

func (m *MathMock) Plus(a int, b int) int {
	return 0
}

func TestSum(t *testing.T) {
	math := &Math{}
	assert.Equal(t, 4, math.Plus(2, 2))
}

func TestDividir(t *testing.T) {
	math := &Math{}
	result, err := math.Divide(4, 2)

	assert.Equal(t, nil, err)
	assert.Equal(t, 2, result)
}

func TestDivideByZero(t *testing.T) {
	math := &Math{}
	result, err := math.Divide(4, 0)

	assert.Equal(t, 0, result)
	assert.NotNil(t, err)
}
