package iban

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIbanValidation(t *testing.T) {

	t.Run("iban number should be valid", func(t *testing.T) {
		isValid, _ := isValidIbanNum("GB33BUKB20201555555555")
		assert.Equal(t, isValid, true)
	})

	t.Run("iban number should be valid for lower case", func(t *testing.T) {
		isValid, _ := isValidIbanNum("gb33bukb20201555555555")
		assert.Equal(t, isValid, true)
	})

	t.Run("iban number should not be valid when structure is wrong", func(t *testing.T) {
		isValid, reason := isValidIbanNum("GB33BUKB2020155555555A")
		assert.Equal(t, isValid, false)
		assert.Equal(t, reason, BBANInvalid)
	})

	t.Run("iban number should not be valid when check digit is wrong", func(t *testing.T) {
		isValid, reason := isValidIbanNum("GB34BUKB20201555555555")
		assert.Equal(t, isValid, false)
		assert.Equal(t, reason, checkDigitInvalid)
	})

	t.Run("iban number should not be valid when it is too short", func(t *testing.T) {
		isValid, reason := isValidIbanNum("GB33BUKB202")
		assert.Equal(t, isValid, false)
		assert.Equal(t, reason, lengthInvalid)
	})

	t.Run("iban number should not be valid when it is empty", func(t *testing.T) {
		isValid, reason := isValidIbanNum("")
		assert.Equal(t, isValid, false)
		assert.Equal(t, reason, lengthInvalid)
	})

	t.Run("iban number should not be valid when country is not available", func(t *testing.T) {
		isValid, reason := isValidIbanNum("ZN33BUKB20201555555555")
		assert.Equal(t, isValid, false)
		assert.Equal(t, reason, countryNotSupport)
	})
}
