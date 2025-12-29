package iteration

import "strings"

// slower, using add and assignment operator
// func Repeat(char string, iterations int) string {
// 	var repeated string
// 	for i := 0; i < iterations; i++ {
// 		repeated += char
// 	}
// 	return repeated
// }

// faster, using strings.Builder
// func Repeat(char string, iterations int) string {
// 	var repeated strings.Builder
// 	for i := 0; i < iterations; i++ {
// 		repeated.WriteString(char)
// 	}
// 	return repeated.String()
// }

// "cleaner", using strings.Repeat
func Repeat(char string, iterations int) string {
	return strings.Repeat(char, iterations)
}
