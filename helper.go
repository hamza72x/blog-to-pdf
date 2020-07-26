package main

import "strings"

// $ => #
func hashifyDollar(str string) string {
	return strings.ReplaceAll(str, "$", "#")
}

// $ => ;
func collonifyDollar(str string) string {
	return strings.ReplaceAll(str, "$", ";")
}

// [$] => [#]
func hashifyDollars(strs []string) []string {
	for i := range strs {
		strs[i] = hashifyDollar(strs[i])
	}
	return strs
}
