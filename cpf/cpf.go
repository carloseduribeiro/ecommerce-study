package cpf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	factorDigit1 = 10
	factorDigit2 = 11
)

type CPF string

func (c CPF) Validate() bool {
	cpf := cleanCpf(c)
	if !isValidLength(cpf) {
		return false
	}
	if hasAllDigitsEquals(cpf) {
		return false
	}
	digit1 := calculateCheckDigit(cpf, factorDigit1)
	digit2 := calculateCheckDigit(cpf, factorDigit2)
	checkDigit := extractCheckDigit(cpf)
	calculatedCheckDigit := fmt.Sprintf("%d%d", digit1, digit2)
	return checkDigit == calculatedCheckDigit
}

func extractCheckDigit(cpf string) string {
	return cpf[len(cpf)-2:]
}

func calculateCheckDigit(cpf string, factor int) int {
	var total int
	for _, char := range strings.Split(cpf, "") {
		if factor > 1 {
			digit, _ := strconv.Atoi(char)
			total += digit * factor
			factor--
		}
	}
	rest := total % 11
	if rest >= 2 {
		return 11 - rest
	}
	return 0
}

func hasAllDigitsEquals(cpf string) bool {
	firstDigit := string(cpf[0])
	return strings.Count(cpf, firstDigit) == len(cpf)
}

func isValidLength[T string | CPF](cpf T) bool {
	if len(cpf) == 11 {
		return true
	}
	return false
}

func cleanCpf[T string | CPF](cpf T) string {
	reg, _ := regexp.Compile("\\d+")
	result := reg.FindAllString(string(cpf), -1)
	return strings.Join(result, "")
}
