package main

import (
	"fmt"

	"github.com/elisiei/re"
)

// thanks chatgpt for these examples, i am too lazy.
func main() {
	// Basic exact match
	exact := re.Exactly("hello")
	reg := re.Create(exact)
	fmt.Println("Exact:", reg.String())
	fmt.Println("Match 'hello':", reg.MatchString("hello"))

	// AnyOf / alternatives
	alt := re.AnyOf("foo", "bar", "baz")
	fmt.Println("AnyOf:", alt.String())
	reg = re.Create(alt)
	fmt.Println("Match 'bar':", reg.MatchString("bar"))

	// Optional, OneOrMore, ZeroOrMore
	opt := re.Exactly("a").Optionally()
	one := re.Exactly("b").OneOrMore()
	zero := re.Exactly("c").ZeroOrMore()
	reg = re.Create(opt, one, zero)
	fmt.Println("Optional + OneOrMore + ZeroOrMore:", reg.String())
	fmt.Println("Match 'abcc':", reg.MatchString("abcc"))
	fmt.Println("Match 'b':", reg.MatchString("b"))

	// Character classes
	chars := re.CharIn("abc").OneOrMore()
	reg = re.Create(chars)
	fmt.Println("CharIn:", reg.String())
	fmt.Println("Match 'aabbc':", reg.MatchString("aabbc"))

	charsNot := re.CharNotIn("xyz").OneOrMore()
	reg = re.Create(charsNot)
	fmt.Println("CharNotIn:", reg.String())
	fmt.Println("Match 'abc':", reg.MatchString("abc"))
	fmt.Println("Match 'xyz':", reg.MatchString("xyz"))

	// Predefined helpers
	helper := re.Digit().OneOrMore().As("digits")
	reg = re.Create(helper)
	fmt.Println("Digit OneOrMore As:", reg.String())
	text := "My number is 12345"
	match := reg.FindStringSubmatch(text)
	if match != nil {
		names := reg.SubexpNames()
		for i, name := range names {
			if i == 0 || name == "" {
				continue
			}
			fmt.Printf("%s = %s\n", name, match[i])
		}
	}

	// Word, WordChar, WordBoundary, Whitespace
	pattern := re.Word().And(re.Whitespace()).And(re.WordChar())
	reg = re.Create(pattern)
	fmt.Println("Word + Whitespace + WordChar:", reg.String())
	fmt.Println("Match 'hello x':", reg.MatchString("hello x"))

	// Grouped, nested groups
	grouped := re.Exactly("foo").Grouped().Or(re.Exactly("bar"))
	reg = re.Create(grouped)
	fmt.Println("Grouped Or:", reg.String())
	fmt.Println("Match 'bar':", reg.MatchString("bar"))

	// Times
	times := re.Exactly("ha").Times(3)
	reg = re.Create(times)
	fmt.Println("Times 3:", reg.String())
	fmt.Println("Match 'hahaha':", reg.MatchString("hahaha"))

	// Line start/end
	line := re.Exactly("start").AtLineStart().And(re.Exactly("end").AtLineEnd())
	reg = re.Create(line)
	fmt.Println("Line start/end:", reg.String())
	fmt.Println("Match 'startend':", reg.MatchString("startend"))
	fmt.Println("Match 'start end':", reg.MatchString("start end"))

	// Or chaining with Maybe
	complex := re.Exactly("foo").And(re.Maybe("bar")).Or(re.Exactly("baz"))
	reg = re.Create(complex)
	fmt.Println("Complex pattern:", reg.String())
	fmt.Println("Match 'foobar':", reg.MatchString("foobar"))
	fmt.Println("Match 'foo':", reg.MatchString("foo"))
	fmt.Println("Match 'baz':", reg.MatchString("baz"))

	// Full example with multiple named captures
	full := re.Digit().OneOrMore().As("major").
		And(re.Exactly(".").And(re.Digit().OneOrMore().As("minor"))).
		And(re.Exactly(".").And(re.Digit().OneOrMore().As("patch")).Optionally())
	reg = re.Create(full)
	fmt.Println("Semver pattern:", reg.String())
	text = "Version 2.10.3"
	m := reg.FindStringSubmatch(text)
	if m != nil {
		names := reg.SubexpNames()
		for i, name := range names {
			if i == 0 || name == "" {
				continue
			}
			fmt.Printf("%s = %s\n", name, m[i])
		}
	}
}
