package packages

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	. "github.com/stretchr/testify/assert"
)

// https://github.com/TannerGabriel/learning-go/blob/master/basics/17-Regex/Regex.go
func TestAllRegex(t *testing.T) {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	True(t, match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	True(t, r.Match([]byte("peach")))
	True(t, r.MatchString("peach"))
	Equal(t,
		"peach",
		r.FindString("peach punch"),
	)
	Equal(t,
		[][]byte{
			[]byte("peach"),
			[]byte("punch"),
		},
		r.FindAll([]byte(`mypeach punch`), -1),
	)
	Equal(t,
		[]int{0, 5},
		r.FindStringIndex("peach punch"),
	)
	Equal(t,
		[]string{"peach", "ea"},
		r.FindStringSubmatch("peach punch"),
	)
	Equal(t,
		[]int{0, 5, 1, 3},
		r.FindStringSubmatchIndex("peach punch"),
	)
	Equal(t,
		[]string{"peach", "punch", "pinch"},
		r.FindAllString("peach punch pinch", -1),
	)
	Equal(t,
		[]string{"peach", "punch"},
		r.FindAllString("peach punch pinch", 2),
	)
	Equal(t,
		[][]int{{0, 5, 1, 3}, {6, 11, 7, 9}, {12, 17, 13, 15}},
		r.FindAllStringSubmatchIndex("peach punch pinch", -1),
	)

	r = regexp.MustCompile("p([a-z]+)ch")
	Equal(t,
		"p([a-z]+)ch",
		r.String(),
	)
	Equal(t,
		"a <fruit>",
		r.ReplaceAllString("a peach", "<fruit>"),
	)
	in := []byte("a peach")
	Equal(t,
		[]byte("a PEACH"),
		r.ReplaceAllFunc(in, bytes.ToUpper),
	)
}

// * 是否匹配 bool
func TestIsMatchWithBool(t *testing.T) {
	var re = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	True(t, re.MatchString("adam[23]"))
	True(t, re.MatchString("eve[7]"))
	False(t, re.MatchString("Job[48]"))
	False(t, re.MatchString("snakey"))
}

// * 是否匹配 with error. Match (Contains any words), This also works for MatchString and MatchReader, just the param are different.
func TestIsMatchWithBoolAndError(t *testing.T) {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	True(t, matched)
	Nil(t, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	False(t, matched)
	Nil(t, err)
	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	False(t, matched)
	NotNil(t, err) // error parsing regexp: missing closing ): `a(b`
}

func TestFind(t *testing.T) {
	// * Example 1
	re := regexp.MustCompile(`foo.?`)
	Equal(t, "food", re.FindString("seafood fool"))
	Equal(t, "", re.FindString("meat"))
	// * Example 2
	re = regexp.MustCompile(`a.`)
	Equal(t,
		[]string{"ar", "an", "al"},
		re.FindAllString("paranormal", -1),
	)
	Equal(t,
		[]string{"ar", "an"},
		re.FindAllString("paranormal", 2),
	)
	Equal(t,
		[]string{"aa"},
		re.FindAllString("graal", -1),
	)
	Nil(t, re.FindAllString("none", -1))

	// * Example 3
	str := "Hello @world@ Match"
	re, _ = regexp.Compile("@(.*)@")
	m := re.FindStringSubmatch(str)
	if len(m) > 1 {
		Equal(t, m[1], "world")
	}

	// * Example 4:
	re = regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1)) // [["food" "d"] ["fool" "l"]]

	// * Find index of the word.
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	loc := pattern.FindIndex(content) // This also works for FindStringIndex.
	NotNil(t, loc)
	Equal(t,
		[]int{18, 33},
		loc,
	)
	Equal(t,
		"option1: value1",
		string(content[loc[0]:loc[1]]),
	)

	// * Find index of the submatched word.
	re = regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	Equal(t,
		[]int{1, 3, 2, 2},
		re.FindSubmatchIndex([]byte("-ab-")),
	)
	Equal(t,
		[]int{1, 5, 2, 4},
		re.FindSubmatchIndex([]byte("-axxb-")),
	)
	Equal(t,
		[]int{1, 3, 2, 2},
		re.FindSubmatchIndex([]byte("-ab-axb-")),
	)
	Equal(t,
		[]int{1, 5, 2, 4},
		re.FindSubmatchIndex([]byte("-axxb-ab-")),
	)
	Nil(t, re.FindSubmatchIndex([]byte("-foo-")))

	// * Longest match
	re = regexp.MustCompile(`a(|b)`)
	Equal(t, "a", re.FindString("ab"))
	re.Longest()
	Equal(t, "ab", re.FindString("ab"))

	// * Find index of the word.
	content = []byte("London")
	re = regexp.MustCompile(`o.`)
	Equal(t,
		[][]int{{1, 3}},
		re.FindAllIndex(content, 1),
	)
	Equal(t,
		[][]int{{1, 3}, {4, 6}},
		re.FindAllIndex(content, -1),
	)

	// * Find index of the word by the name.
	// By `?P<yourNames>`
	re = regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	matches := re.FindStringSubmatch("Alan Turing")
	lastIndex := re.SubexpIndex("last") // Get the index by the name.
	Equal(t, 2, lastIndex)
	Equal(t, "Turing", matches[lastIndex])
}

// * Replace content.
func TestReplaceContent(t *testing.T) {
	// * Example 1: With Expand.
	content := []byte(`
			# comment line
			option1: value1
			option2: value2

			# another comment line
			option3: value3
		`)
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	template := []byte("$key=$value\n")
	res := []byte{}
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		res = pattern.Expand(res, template, content, submatches)
	}
	ptr(string(res))

	// * Example 2: With SubexpNames().
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	Equal(t,
		[]string{"", "first", "last"},
		re.SubexpNames(),
	)
	reversed := fmt.Sprintf("${%s} ${%s}",
		re.SubexpNames()[2], re.SubexpNames()[1])
	Equal(t,
		"${last} ${first}",
		reversed,
	)
	Equal(t,
		"Turing Alan",
		re.ReplaceAllString("Alan Turing", reversed),
	)
}

func TestSplitStringByRegexp(t *testing.T) {
	re := regexp.MustCompile(`a`)
	Nil(t, re.Split("banana", 0))
	Equal(t,
		[]string{"b", "n", "n", ""},
		re.Split("banana", -1),
	)
	Equal(t,
		[]string{"banana"},
		re.Split("banana", 1),
	)
	Equal(t,
		[]string{"b", "nana"},
		re.Split("banana", 2),
	)

	re = regexp.MustCompile(`z+`)
	Nil(t, re.Split("pizza", 0))
	Equal(t,
		[]string{"pi", "a"},
		re.Split("pizza", -1),
	)
	Equal(t,
		[]string{"pizza"},
		re.Split("pizza", 1),
	)
	Equal(t,
		[]string{"pi", "a"},
		re.Split("pizza", 2),
	)
}

func TestMatchNumberExample(t *testing.T) {
	// * Example: matching a number.
	pattern := "[0-9]+"
	re, _ := regexp.Compile(pattern)
	str := "Some string 0 1 2 3 4 "

	// * Match condition.
	True(t, re.MatchString(str))

	// * Return match
	res := re.FindString(str)
	Equal(t,
		"0",
		res,
	)
	// * Return multiple matches
	results := re.FindAllString(str, -1)
	Equal(t,
		[]string{"0", "1", "2", "3", "4"},
		results,
	)
	// * Replace match (This also works with ReplaceAll)
	replaceResults := re.ReplaceAllString(str, "num")
	Equal(t,
		"Some string num num num num num ",
		replaceResults,
	)
}

func TestVerifyEmail(t *testing.T) {
	isValidEmail := func(email string) bool {
		var emailRegexp = regexp.MustCompile("(?i)" + // Case insensitive
			"^[a-z0-9!#$%&'*+/=?^_`{|}~.-]+" + // Validate local part
			"@" +
			"[a-z0-9-]+(\\.[a-z0-9-]+)*$") // Validate domain part
		if len(email) > 254 {
			return false
		}
		return emailRegexp.MatchString(email)
	}

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
		Equal(t, tC.valid, result)
	}
}
