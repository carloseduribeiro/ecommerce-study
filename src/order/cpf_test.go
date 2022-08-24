package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{name: "Test with valid cpf", cpf: "17185070031", expected: true},
		{name: "Test invalid cpf with different digits ", cpf: "93847575438", expected: false},
		{name: "Test with a CPF with all digits the same", cpf: "99999999999", expected: false},
		{name: "Test a cpf with more than eleven digits", cpf: "121212121212", expected: false},
		{name: "Test a cpf with less than eleven digits", cpf: "1010101010", expected: false},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			obtained := validate(tt.cpf)
			assert.EqualValues(t, tt.expected, obtained)
		})
	}
}

func TestNewCPF(t *testing.T) {
	t.Run("should create a new CPF", func(t *testing.T) {
		value := "17185070031"
		expected := &CPF{value: value}
		obtained, err := NewCPF(value)
		assert.NoError(t, err)
		assert.Equal(t, expected.Value(), obtained.Value())
	})

	t.Run("should return an error if the value is invalid", func(t *testing.T) {
		value := "93847575438"
		_, err := NewCPF(value)
		assert.Error(t, err)
	})
}
