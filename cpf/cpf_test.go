package cpf

import (
	"github.com/ecommerce-study/tests/assertions"
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name     string
		cpf      CPF
		expected bool
	}{
		{name: "Test with valid cpf", cpf: CPF("19890600560"), expected: true},
		{name: "Test invalid cpf with different digits ", cpf: CPF("93847575438"), expected: false},
		{name: "Test with a CPF with all digits the same", cpf: CPF("99999999999"), expected: false},
		{name: "Test a cpf with more than eleven digits", cpf: CPF("121212121212"), expected: false},
		{name: "Test a cpf with less than eleven digits", cpf: CPF("1010101010"), expected: false},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			obtained := tt.cpf.Validate()
			assertions.AssertEquals(t, tt.expected, obtained)
		})
	}
}
