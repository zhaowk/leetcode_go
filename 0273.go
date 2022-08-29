package leetcode_go

import "fmt"

const (
	Zero = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Eleven
	Twelve
	Thirteen
	Fourteen
	Fifteen
	Sixteen
	Seventeen
	Eighteen
	Nineteen
	Twenty
	Thirty = 30
	Forty = 40
	Fifty = 50
	Sixty = 60
	Seventy = 70
	Eighty = 80
	Ninety = 90
	Hundred = 100
	Thousand = 1000
	Million = 1000000
	Billion = 1000000000
)

var (
	numsMap = map[int]string {
		Zero: "Zero",
		One : "One",
		Two : "Two",
		Three : "Three",
		Four : "Four",
		Five : "Five",
		Six : "Six",
		Seven : "Seven",
		Eight : "Eight",
		Nine : "Nine",
		Ten : "Ten",
		Eleven : "Eleven",
		Twelve : "Twelve",
		Thirteen : "Thirteen",
		Fourteen : "Fourteen",
		Fifteen : "Fifteen",
		Sixteen : "Sixteen",
		Seventeen : "Seventeen",
		Eighteen : "Eighteen",
		Nineteen : "Nineteen",
		Twenty : "Twenty",
		Thirty : "Thirty",
		Forty : "Forty",
		Fifty : "Fifty",
		Sixty : "Sixty",
		Seventy : "Seventy",
		Eighty : "Eighty",
		Ninety : "Ninety",
		Hundred : "Hundred",
		Thousand : "Thousand",
		Million : "Million",
		Billion : "Billion",
	}
)

func numToWordsHundreds(num int) string {
	result := ""
	space := ""  // space after Hundreds
	h := num / 100

	if h != 0 {
		result = fmt.Sprintf("%s %s", numsMap[h], numsMap[Hundred])
		space = result + " "
	}

	h = num % 100

	switch {
	case 1 <= h && h <= 20:
		result = fmt.Sprintf("%s%s", space, numsMap[h])
	case 21 <= h && h <= 99 :
		if h % 10 == 0 {
			result = fmt.Sprintf("%s%s", space, numsMap[h])
		} else {
			result = fmt.Sprintf("%s%s %s", space, numsMap[h / 10 * 10], numsMap[h % 10])
		}
	case h == 0:  // do nothing
	default: // impossible
	}
	return result
}

func numberToWords(num int) string {
	if num == 0 {
		return numsMap[Zero]
	}

	result := numToWordsHundreds(num % 1000)
	unit := ""

	num /= 1000
	unit = numsMap[Thousand]

	if num > 0 && num % 1000 != 0 {
		t := ""
		if len(result) > 0 {
			t = " " + result
		}
		result = fmt.Sprintf("%s %s%s", numToWordsHundreds(num % 1000), unit, t)
	}

	num /= 1000
	unit = numsMap[Million]
	if num > 0 && num % 1000 != 0 {
		t := ""
		if len(result) > 0 {
			t = " " + result
		}
		result = fmt.Sprintf("%s %s%s", numToWordsHundreds(num % 1000), unit, t)
	}

	num /= 1000
	unit = numsMap[Billion]
	if num > 0 && num % 1000 != 0 {
		t := ""
		if len(result) > 0 {
			t = " " + result
		}
		result = fmt.Sprintf("%s %s%s", numToWordsHundreds(num % 1000), unit, t)
	}

	return result
}
