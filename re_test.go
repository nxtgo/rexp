package re_test

import (
	"regexp"
	"testing"

	"github.com/elisiei/re"
)

func TestReHelpers(t *testing.T) {
	tests := []struct {
		name       string
		pattern    *regexp.Regexp
		input      string
		wantMatch  bool
		wantGroups map[string]string
	}{
		// Exactly
		{"Exactly match", re.Create(re.Exactly("hello")), "hello", true, nil},
		{"Exactly no match", re.Create(re.Exactly("hello")), "hell", false, nil},

		// AnyOf
		{"AnyOf match", re.Create(re.AnyOf("foo", "bar", "baz")), "bar", true, nil},
		{"AnyOf no match", re.Create(re.AnyOf("foo", "bar", "baz")), "qux", false, nil},

		// Optional / Maybe
		{"Optional match present", re.Create(re.Exactly("a").Optionally()), "a", true, nil},
		{"Optional match absent", re.Create(re.Exactly("a").Optionally()), "", true, nil},

		// OneOrMore / ZeroOrMore
		{"OneOrMore match", re.Create(re.Exactly("b").OneOrMore()), "bbb", true, nil},
		{"ZeroOrMore match empty", re.Create(re.Exactly("c").ZeroOrMore()), "", true, nil},

		// CharIn / CharNotIn
		{"CharIn match", re.Create(re.CharIn("abc").OneOrMore()), "aabbc", true, nil},
		{"CharNotIn match", re.Create(re.CharNotIn("xyz").OneOrMore()), "abc", true, nil},

		// Word / WordChar / Whitespace
		{"Word match", re.Create(re.Word()), "hello", true, nil},
		{"WordChar match", re.Create(re.WordChar()), "x", true, nil},
		{"Whitespace match", re.Create(re.Whitespace()), " ", true, nil},

		// Grouped / As (named capture)
		{"Named capture", re.Create(re.Digit().OneOrMore().As("num")), "123", true, map[string]string{"num": "123"}},

		// And / Or chaining
		{"And match", re.Create(re.Exactly("foo").And(re.Exactly("bar"))), "foobar", true, nil},
		{"Or match", re.Create(re.Exactly("foo").Or(re.Exactly("bar"))), "bar", true, nil},

		// Times
		{"Times match", re.Create(re.Exactly("ha").Times(3)), "hahaha", true, nil},

		// Line anchors
		{"AtLineStart match", re.Create(re.Exactly("start").AtLineStart()), "start", true, nil},
		{"AtLineEnd match", re.Create(re.Exactly("end").AtLineEnd()), "end", true, nil},

		// Complex semver example with named captures
		{"Semver capture", re.Create(
			re.Digit().OneOrMore().As("major").
				And(re.Exactly(".").And(re.Digit().OneOrMore().As("minor"))).
				And(re.Exactly(".").And(re.Digit().OneOrMore().As("patch")).Optionally()),
		), "2.10.3", true, map[string]string{"major": "2", "minor": "10", "patch": "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pattern.MatchString(tt.input)
			if got != tt.wantMatch {
				t.Errorf("MatchString(%q) = %v; want %v", tt.input, got, tt.wantMatch)
			}

			if tt.wantGroups != nil {
				matches := tt.pattern.FindStringSubmatch(tt.input)
				if matches == nil {
					t.Fatalf("Expected match but got nil")
				}
				names := tt.pattern.SubexpNames()
				for i, name := range names {
					if name == "" {
						continue
					}
					if matches[i] != tt.wantGroups[name] {
						t.Errorf("Group %q = %q; want %q", name, matches[i], tt.wantGroups[name])
					}
				}
			}
		})
	}
}
