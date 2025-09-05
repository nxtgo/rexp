package rexp_test

import (
	"regexp"
	"testing"

	"github.com/nxtgo/rexp"
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
		{"Exactly match", rexp.Create(rexp.Exactly("hello")), "hello", true, nil},
		{"Exactly no match", rexp.Create(rexp.Exactly("hello")), "hell", false, nil},

		// AnyOf
		{"AnyOf match", rexp.Create(rexp.AnyOf("foo", "bar", "baz")), "bar", true, nil},
		{"AnyOf no match", rexp.Create(rexp.AnyOf("foo", "bar", "baz")), "qux", false, nil},

		// Optional / Maybe
		{"Optional match present", rexp.Create(rexp.Exactly("a").Optionally()), "a", true, nil},
		{"Optional match absent", rexp.Create(rexp.Exactly("a").Optionally()), "", true, nil},

		// OneOrMore / ZeroOrMore
		{"OneOrMore match", rexp.Create(rexp.Exactly("b").OneOrMore()), "bbb", true, nil},
		{"ZeroOrMore match empty", rexp.Create(rexp.Exactly("c").ZeroOrMore()), "", true, nil},

		// CharIn / CharNotIn
		{"CharIn match", rexp.Create(rexp.CharIn("abc").OneOrMore()), "aabbc", true, nil},
		{"CharNotIn match", rexp.Create(rexp.CharNotIn("xyz").OneOrMore()), "abc", true, nil},

		// Word / WordChar / Whitespace
		{"Word match", rexp.Create(rexp.Word()), "hello", true, nil},
		{"WordChar match", rexp.Create(rexp.WordChar()), "x", true, nil},
		{"Whitespace match", rexp.Create(rexp.Whitespace()), " ", true, nil},

		// Grouped / As (named capture)
		{"Named capture", rexp.Create(rexp.Digit().OneOrMore().As("num")), "123", true, map[string]string{"num": "123"}},

		// And / Or chaining
		{"And match", rexp.Create(rexp.Exactly("foo").And(rexp.Exactly("bar"))), "foobar", true, nil},
		{"Or match", rexp.Create(rexp.Exactly("foo").Or(rexp.Exactly("bar"))), "bar", true, nil},

		// Times
		{"Times match", rexp.Create(rexp.Exactly("ha").Times(3)), "hahaha", true, nil},

		// Line anchors
		{"AtLineStart match", rexp.Create(rexp.Exactly("start").AtLineStart()), "start", true, nil},
		{"AtLineEnd match", rexp.Create(rexp.Exactly("end").AtLineEnd()), "end", true, nil},

		// Complex semver example with named captures
		{"Semver capture", rexp.Create(
			rexp.Digit().OneOrMore().As("major").
				And(rexp.Exactly(".").And(rexp.Digit().OneOrMore().As("minor"))).
				And(rexp.Exactly(".").And(rexp.Digit().OneOrMore().As("patch")).Optionally()),
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
