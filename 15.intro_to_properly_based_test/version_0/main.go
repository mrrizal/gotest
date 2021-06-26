package main

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(input int) string {
	var result strings.Builder

	for _, romanNumber := range allRomanNumerals {
		for input >= romanNumber.Value {
			input -= romanNumber.Value
			result.WriteString(romanNumber.Symbol)
		}
	}

	return result.String()
}

var romanNumbers = make(map[string]int)

func initRomanNumbers() {
	romanNumbers["M"] = 1000
	romanNumbers["CM"] = 900
	romanNumbers["D"] = 500
	romanNumbers["CD"] = 400
	romanNumbers["C"] = 100
	romanNumbers["XC"] = 90
	romanNumbers["L"] = 50
	romanNumbers["XL"] = 40
	romanNumbers["X"] = 10
	romanNumbers["IX"] = 9
	romanNumbers["V"] = 5
	romanNumbers["IV"] = 4
	romanNumbers["I"] = 1
}

func parseRomanNumber(roman *string) string {
	var tempRomanNumber string
	var n int
	if len(*roman) > 1 {
		n = 2
		tempRomanNumber = (*roman)[0:n]
		if _, ok := romanNumbers[tempRomanNumber]; !ok {
			n = 1
			tempRomanNumber = (*roman)[0:n]
		}
	} else {
		n = 1
		tempRomanNumber = (*roman)[0:n]
	}

	*roman = (*roman)[n:]
	return tempRomanNumber
}

func ConvertToArabic(roman string) int {
	initRomanNumbers()
	arabic := 0
	n := len(roman)
	for i := 0; i < n; i++ {
		tempRomanNumber := parseRomanNumber(&roman)
		arabic += romanNumbers[tempRomanNumber]
		if roman == "" {
			break
		}
	}
	return arabic
}

func main() {
	roman := "MMMCMXCIX"
	fmt.Println(ConvertToArabic(roman))
}
