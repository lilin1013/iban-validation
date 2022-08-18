package iban

import (
	"fmt"
	"strings"
)

const countryNotSupport = "countryNotSupport"
const lengthInvalid = "lengthInvalid"
const BBANInvalid = "BBANInvalid"
const checkDigitInvalid = "checkDigitInvalid"

type Iban struct {
	CountryCode    string
	CheckDigit     string
	BBAN           string
	length         int
	countrySetting CountrySetting
}

func isValidIbanNum(ibanNumber string) (bool, string) {
	ibanNumber = strings.ToUpper(strings.Replace(ibanNumber, " ", "", -1))

	length := len(ibanNumber)
	if length < 16 {
		return false, lengthInvalid
	}

	countryCode := ibanNumber[0:2]
	checkDigit := ibanNumber[2:4]
	bban := ibanNumber[4:]

	countrySetting, err := getCountrySetting(countryCode)
	if err != nil {
		return false, countryNotSupport
	}

	ibanItem := Iban{CheckDigit: checkDigit, CountryCode: countryCode, BBAN: bban, length: length, countrySetting: countrySetting}

	return ibanItem.isValidIban()
}

func (iban *Iban) isValidIban() (bool, string) {

	err := iban.isValidIBANLength()
	if err != nil {
		return false, lengthInvalid
	}

	err = iban.isValidBBAN()
	if err != nil {
		return false, BBANInvalid
	}

	err = iban.isValidIBANCheckDigit()
	if err != nil {
		return false, checkDigitInvalid
	}

	return true, ""
}

func (iban *Iban) isValidIBANLength() error {
	if iban.length != iban.countrySetting.Length {
		return fmt.Errorf("the length is not match")
	}
	return nil
}

func (iban *Iban) isValidBBAN() error {
	return nil
}

func (iban *Iban) isValidIBANCheckDigit() error {
	return nil
}

func getCountrySetting(countryCode string) (CountrySetting, error) {
	if item, ok := countries[countryCode]; ok {
		return item, nil
	}
	return CountrySetting{}, fmt.Errorf("Do not support country: %v", countryCode)
}
