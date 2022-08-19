package iban

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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
	bban := iban.BBAN
	format := iban.countrySetting.Format

	//format quantifier, e.g. 08 -> {8}
	formatter := regexp.MustCompile(`\d{2}`)
	format = formatter.ReplaceAllStringFunc(format, func(m string) string {
		quantifier, atoiErr := strconv.Atoi(m)
		if atoiErr != nil {
			return ""
		}

		return fmt.Sprintf("{%d}", quantifier)
	})

	//format the character set
	formatter = regexp.MustCompile(`[AFU]`)
	format = formatter.ReplaceAllStringFunc(format, func(m string) string {
		return countryReg[m]
	})

	regexFormatter, err := regexp.Compile(format)
	if err != nil {
		return fmt.Errorf("failed compile the regexp: %v", err.Error())
	}

	if !regexFormatter.MatchString(bban) {
		return errors.New("BBAN part of IBAN is not formatted according to country specification")
	}

	return nil
}

func (iban *Iban) isValidIBANCheckDigit() error {
	//rearrange iban
	newIban := iban.BBAN + iban.CountryCode + iban.CheckDigit

	modStr := replaceCharToInt(newIban)

	if mod97(modStr) != 1 {
		return errors.New("not valid check digit")
	}

	return nil
}

func getCountrySetting(countryCode string) (CountrySetting, error) {
	if item, ok := countries[countryCode]; ok {
		return item, nil
	}
	return CountrySetting{}, fmt.Errorf("Do not support country: %v", countryCode)
}

func replaceCharToInt(str string) string {
	// replace the char A - Z to int
	formatParts := regexp.MustCompile(`[A-Z]`)
	return formatParts.ReplaceAllStringFunc(str, func(m string) string {
		return strconv.Itoa(int(m[0] - 55))
	})
}

func mod97(str string) int {

	// initial remaining is 0 and remaining string is ""
	resStr := ""
	res := 0

	//the loop to do mod97 algorithm until no digits in the str left
	for ok := true; ok; ok = len(str) > 0 {

		//construct remaining String
		if res > 0 {
			resStr = strconv.Itoa(res)
		} else {
			resStr = ""
		}

		//calculate how many digits needed to construct 9 digits, besides the previous remaining digit
		n := 9 - len(resStr)

		//if less digit than needed, construct with remaining digit
		if len(str) < n {
			n = len(str)
		}

		//construct num String to be mod
		newStr := resStr + str[:n]

		value, err := strconv.Atoi(newStr)
		if err != nil {
			return 0
		}

		// mod by 97
		res = value % 97

		//redefine the digits left, prepare for the next calculation
		str = str[n:]
	}

	return res
}
