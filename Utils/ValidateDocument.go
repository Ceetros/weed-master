package Utils

import (
	"strconv"
	"strings"
)

func ValidateDocument(doc string) bool {
	doc = onlyNumbers(doc)

	if len(doc) == 11 {
		return validateCPF(doc)
	}
	if len(doc) == 14 {
		return validateCNPJ(doc)
	}
	return false
}

// Valida CPF
func validateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	for i := 0; i < 10; i++ {
		if cpf == strings.Repeat(strconv.Itoa(i), 11) {
			return false
		}
	}

	sum := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (10 - i)
	}
	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}
	if d1 != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (11 - i)
	}
	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}
	return d2 == int(cpf[10]-'0')
}

func validateCNPJ(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}

	for i := 0; i < 10; i++ {
		if cnpj == strings.Repeat(strconv.Itoa(i), 14) {
			return false
		}
	}

	mult1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	mult2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * mult1[i]
	}
	d1 := sum % 11
	if d1 < 2 {
		d1 = 0
	} else {
		d1 = 11 - d1
	}
	if d1 != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * mult2[i]
	}
	d2 := sum % 11
	if d2 < 2 {
		d2 = 0
	} else {
		d2 = 11 - d2
	}
	return d2 == int(cnpj[13]-'0')
}
func onlyNumbers(value string) string {
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, "-", "")
	value = strings.ReplaceAll(value, "/", "")
	return strings.TrimSpace(value)
}
