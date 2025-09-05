package main

import (
	"fmt"

	"github.com/nxtgo/rexp"
)

// thanks chatgpt for these examples, i am too lazy.
func main() {
	// Basic exact match
	exact := rexp.Exactly("hello")
	reg := rexp.Create(exact)
	fmt.Println("Exact:", reg.String())
	fmt.Println("Match 'hello':", reg.MatchString("hello"))

	// AnyOf / alternatives
	alt := rexp.AnyOf("foo", "bar", "baz")
	fmt.Println("AnyOf:", alt.String())
	reg = rexp.Create(alt)
	fmt.Println("Match 'bar':", reg.MatchString("bar"))

	// Optional, OneOrMore, ZeroOrMore
	opt := rexp.Exactly("a").Optionally()
	one := rexp.Exactly("b").OneOrMore()
	zero := rexp.Exactly("c").ZeroOrMore()
	reg = rexp.Create(opt, one, zero)
	fmt.Println("Optional + OneOrMore + ZeroOrMore:", reg.String())
	fmt.Println("Match 'abcc':", reg.MatchString("abcc"))
	fmt.Println("Match 'b':", reg.MatchString("b"))

	// Character classes
	chars := rexp.CharIn("abc").OneOrMore()
	reg = rexp.Create(chars)
	fmt.Println("CharIn:", reg.String())
	fmt.Println("Match 'aabbc':", reg.MatchString("aabbc"))

	charsNot := rexp.CharNotIn("xyz").OneOrMore()
	reg = rexp.Create(charsNot)
	fmt.Println("CharNotIn:", reg.String())
	fmt.Println("Match 'abc':", reg.MatchString("abc"))
	fmt.Println("Match 'xyz':", reg.MatchString("xyz"))

	// Predefined helpers
	helper := rexp.Digit().OneOrMore().As("digits")
	reg = rexp.Create(helper)
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
	pattern := rexp.Word().And(rexp.Whitespace()).And(rexp.WordChar())
	reg = rexp.Create(pattern)
	fmt.Println("Word + Whitespace + WordChar:", reg.String())
	fmt.Println("Match 'hello x':", reg.MatchString("hello x"))

	// Grouped, nested groups
	grouped := rexp.Exactly("foo").Grouped().Or(rexp.Exactly("bar"))
	reg = rexp.Create(grouped)
	fmt.Println("Grouped Or:", reg.String())
	fmt.Println("Match 'bar':", reg.MatchString("bar"))

	// Times
	times := rexp.Exactly("ha").Times(3)
	reg = rexp.Create(times)
	fmt.Println("Times 3:", reg.String())
	fmt.Println("Match 'hahaha':", reg.MatchString("hahaha"))

	// Line start/end
	line := rexp.Exactly("start").AtLineStart().And(rexp.Exactly("end").AtLineEnd())
	reg = rexp.Create(line)
	fmt.Println("Line start/end:", reg.String())
	fmt.Println("Match 'startend':", reg.MatchString("startend"))
	fmt.Println("Match 'start end':", reg.MatchString("start end"))

	// Or chaining with Maybe
	complex := rexp.Exactly("foo").And(rexp.Maybe("bar")).Or(rexp.Exactly("baz"))
	reg = rexp.Create(complex)
	fmt.Println("Complex pattern:", reg.String())
	fmt.Println("Match 'foobar':", reg.MatchString("foobar"))
	fmt.Println("Match 'foo':", reg.MatchString("foo"))
	fmt.Println("Match 'baz':", reg.MatchString("baz"))

	// Full example with multiple named captures
	full := rexp.Digit().OneOrMore().As("major").
		And(rexp.Exactly(".").And(rexp.Digit().OneOrMore().As("minor"))).
		And(rexp.Exactly(".").And(rexp.Digit().OneOrMore().As("patch")).Optionally())
	reg = rexp.Create(full)
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
