package packages

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestVerifyEmail(t *testing.T) {
	testCases := []struct {
		email string
		valid bool
	}{
		{"a@test.com", true},
		{"postmaster@test.com", true},
		{"president@kremlin.gov.ru", true},
		{"test@test.co.uk", true},
		{"", false},
		{"test", false},
		{"test.com", false},
		{".com", false},
		{"адрес@пример.рф", false},
		{" space_before@test.com", false},
		{"space between@test.com", false},
		{"\nnewlinebefore@test.com", false},
		{"newline\nbetween@test.com", false},
		{"test@test.com.", false},
		{"asyouallcanseethisemailaddressexceedsthemaximumnumberofcharactersallowedtobeintheemailaddresswhichisnomorethatn254accordingtovariousrfcokaycanistopnowornotyetnoineedmorecharacterstoadd@i.really.cannot.thinkof.what.else.to.put.into.this.invalid.address.net", false},
	}
	for _, tC := range testCases {
		result := isValidEmail(tC.email)
		assert.Equal(t, tC.valid, result)
	}
}

func isValidEmail(email string) bool {
	var emailRegexp = regexp.MustCompile("(?i)" + // Case insensitive
		"^[a-z0-9!#$%&'*+/=?^_`{|}~.-]+" + // Validate local part
		"@" +
		"[a-z0-9-]+(\\.[a-z0-9-]+)*$") // Validate domain part
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}
