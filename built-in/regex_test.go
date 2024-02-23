package packages

import (
	"regexp"
	"testing"
)

// https://github.com/TannerGabriel/learning-go/blob/master/basics/17-Regex/Regex.go
func TestRegex(t *testing.T) {
	// Basic Regexp for matching a number
	pattern := "[0-9]+"

	re, _ := regexp.Compile(pattern)

	str := "Some string 0 1 2 3 4 "
	if re.MatchString(str) {
		// matched
	}

	// Return match
	result := re.FindString(str)
	ptr("Number matched:", result)

	// Return multiple matches
	results := re.FindAllString(str, -1)
	ptr("Number matched: ", results)

	// Replace match
	replaceResults := re.ReplaceAllString(str, "num")
	ptr("Result:", replaceResults)

	// Sub-matches
	str1 := "Hello @world@ Match"
	sub_re, _ := regexp.Compile("@(.*)@")

	m := sub_re.FindStringSubmatch(str1)
	if len(m) > 1 {
		ptr(m[1])
	}
}
