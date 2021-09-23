// Copyright (C) 2021  Antonio Lassandro

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.

// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"math"
	"strings"
)

var ToTwenty = [...]string{
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
}

var ToHundred = [...]string{
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

var ToQuintillion = [...]string{
	"Thousand",
	"Million",
	"Billion",
	"Trillion",
	"Quadrillion",
}

func pow(i int, p int) int {
	return int(math.Pow(float64(i), float64(p)))
}

func NumToEnglishWords(n int) []string {
	if n == 0 {
		return []string{"Zero"}
	}

	if n < 20 {
		return []string{ToTwenty[n-1]}
	}

	if n < 100 {
		return []string{ToHundred[n/10-2], NumToEnglish(n % 10)}
	}

	if n < 1000 {
		return []string{ToTwenty[n/100-1], "Hundred", NumToEnglish(n % 100)}
	}

	for i, w := range ToQuintillion {
		p := i + 1
		if n < pow(1000, p+1) {
			pow_p := pow(1000, p)
			return []string{NumToEnglish(n / pow_p), w, NumToEnglish(n % pow_p)}
		}
	}

	return []string{"error"}
}

func NumToEnglish(n int) string {
	return strings.Join(NumToEnglishWords(n), " ")
}
