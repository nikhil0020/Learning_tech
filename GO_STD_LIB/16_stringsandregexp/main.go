package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func ConvertingStrings() {
	product := "Mango"

	fmt.Println(strings.ToLower(product))
	fmt.Println(product)
	fmt.Println(strings.ToUpper(product))
	fmt.Println(strings.ToTitle(product))

	// ToTitle and ToUpper are producing same result for English
	// But they produce different result of other languages
}

func WorkingWithCharacterCase() {
	product := "Kayak"
	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}
}

func InspectingStrings() {
	description := "A boat for one person"
	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
}

func InspectingStringWithCustomFunction() {
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}

	description := "A boat for one person"

	fmt.Println("IndexFunc: ", strings.IndexFunc(description, isLetterB))

}

// Manipulating Strings

func SplittingStrings() {
	description := "A boat for one person"
	splits := strings.Fields(description)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}

	// splitsUsingSplit := strings.Split(description, "o")
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
	splitsN := strings.SplitN(description, " ", 3)
	for _, x := range splitsN {
		fmt.Println("Split >>" + x + "<<")
	}

	doubleSpaced := "This  is  a  double  spaced  string"

	// Below code create empty element in our slice, to deal with that
	// we have to use the Fields method

	// splitsND := strings.Split(doubleSpaced, " ")

	// for _, x := range splitsND {
	// 	fmt.Println("Split >>" + x + "<<")
	// }

	splitsD := strings.Fields(doubleSpaced)

	for _, x := range splitsD {
		fmt.Println("Field >>" + x + "<<")
	}

	//
	// If the string can be split into more strings than has been specified,
	// then the last element in the result slice will be the unsplit remainder of the string.

}

func SplitUsingCustomFunction() {
	description := "This  is  double  spaced"
	splitter := func(r rune) bool {
		return r == ' '
	}
	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}

	// The custom function receives a rune and returns true if that rune should cause the string to split.
	// The FieldsFunc function is smart enough to deal with repeated characters, like the double spaces in Listing
}

func TrimmingString() {
	username := " Alice"
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>"+trimmed+"<<")

	description := "A boat for one person"
	trimmed = strings.Trim(description, "Asno ")
	fmt.Println("Trimmed:", trimmed)

	prefixTrimmed := strings.TrimPrefix(description, "A boat ")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed:", prefixTrimmed)
	fmt.Println("Not trimmed:", wrongPrefix)
}

func TrimmingWithCustomFunction() {
	description := "A boat for one person"
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimmed := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed)
}

func AlteringString() {
	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)
}

func AlteringStringWithMap() {
	text := "It was a boat. A small boat."
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}

		return r
	}
	mapped := strings.Map(mapper, text)

	fmt.Println("Mapped: ", mapped)
}

func ReplacerMethod() {
	text := "It was a boat. A small boat"
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(text)
	// boat will be replaced with kayak
	// small will be replaced with huge
	fmt.Println("Replaced: ", replaced)
}

func BuildingAndGeneratingStrings() {
	text := "It was a boat. A small boat"
	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	fmt.Println("Joined: ", joined)
}

func StringBuilder() {
	text := "It was a boat. A small boat."
	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("String:", builder.String())
}

// Using Regular Expression

func MatchString() {
	description := "A boat for one person"
	number := "123-45-1234"

	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	match, err = regexp.MatchString(`^\d{3}-\d{2}-\d{4}`, number)

	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}
}

func CompilingAndReusingPattern() {
	pattern, compileErr := regexp.Compile("[A-z]oat")
	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
}

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func CompilePatternMethod() {
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
		"=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description, idx))
	}
}

func SplittingStringUsingRegex() {
	pattern := regexp.MustCompile(" |boat|one")
	description := "Kayak. A boat for one person."
	split := pattern.Split(description, -1)

	for _, s := range split {
		if s != "" {
			fmt.Println("Substring: ", s)
		}
	}
}

// Subexpressions allow parts of a regular expression to be accessed,
// which can make it easier to extract substrings from within a matched region
func UsingSubExpressions() {
	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
	description := "Kayak. A boat for one person."
	str := pattern.FindString(description)
	fmt.Println("Match:", str)

	// I can add subexpressions to identify the regions of content that are important within the pattern,
	pattern2 := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	subs := pattern2.FindStringSubmatch(description)
	for _, s := range subs {
		fmt.Println("Match:", s)
	}

	// Subexpressions are denoted with parentheses

}

func UsingNamesSubexpressions() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	subs := pattern.FindStringSubmatch(description)
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}
}

func ReplacingSubStringUsingRegexp() {
	pattern := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)
}

func main() {

	// ConvertingStrings()
	// WorkingWithCharacterCase()
	// InspectingStrings()
	// InspectingStringWithCustomFunction()
	// SplittingStrings()
	// SplitUsingCustomFunction()
	// TrimmingString()
	// TrimmingWithCustomFunction()
	// AlteringString()
	// AlteringStringWithMap()
	// ReplacerMethod()
	// BuildingAndGeneratingStrings()
	// StringBuilder()
	// MatchString()
	// CompilingAndReusingPattern()
	// CompilePatternMethod()
	// SplittingStringUsingRegex()
	// UsingSubExpressions()
	// UsingNamesSubexpressions()
	ReplacingSubStringUsingRegexp()
}
