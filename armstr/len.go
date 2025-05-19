package armstr

import "unicode/utf8"

// Len implements naive way to calculate count of characters
// in given utf-8 string. It will work correctly and better
// than standard Go len(s) but will fail producing correct
// answer for complex combined grapheme clusters like emoji,
// accented letters and som CJK.
//
// Anyway, Unicode is hard and if exact calculation is required
// something like https://github.com/rivo/uniseg must be used.
func Len(s string) int {
	return utf8.RuneCountInString(s)
}
