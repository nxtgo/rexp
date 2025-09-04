package re

import (
	"regexp"
	"strings"
)

type R struct{ s string }

// go's regexp syntax is weird asfuck :skull:
func (r R) String() string        { return r.s }
func (r R) And(xs ...any) R       { return R{s: r.s + concat(xs...)} }
func (r R) Or(xs ...any) R        { return R{s: "(?:" + r.s + "|" + join(xs, "|") + ")"} }
func (r R) Before(xs ...any) R    { return R{s: r.s + "(?=" + concat(xs...) + ")"} }
func (r R) NotBefore(xs ...any) R { return R{s: r.s + "(?!" + concat(xs...) + ")"} }
func (r R) Times(n int) R         { return R{s: "(?:" + r.s + "){" + itoa(n) + "}"} }
func (r R) Optionally() R         { return R{s: "(?:" + r.s + ")?"} }
func (r R) As(name string) R      { return R{s: "(?P<" + name + ">" + r.s + ")"} }
func (r R) Grouped() R            { return R{s: "(" + r.s + ")"} }
func (r R) AtLineStart() R        { return R{s: "^" + r.s} }
func (r R) AtLineEnd() R          { return R{s: r.s + "$"} }
func (r R) OneOrMore() R          { return R{s: "(?:" + r.s + ")+"} }
func (r R) ZeroOrMore() R         { return R{s: "(?:" + r.s + ")*"} }
func (r R) Maybe() R              { return R{s: "(?:" + r.s + ")?"} }

func Create(xs ...any) *regexp.Regexp {
	return regexp.MustCompile(concat(xs...))
}

func Exactly(xs ...any) R    { return R{s: concatEsc(xs...)} }
func AnyOf(xs ...any) R      { return R{s: "(?:" + join(xs, "|") + ")"} }
func Maybe(xs ...any) R      { return R{s: "(?:" + concat(xs...) + ")?"} }
func OneOrMore(xs ...any) R  { return R{s: "(?:" + concat(xs...) + ")+"} }
func ZeroOrMore(xs ...any) R { return R{s: "(?:" + concat(xs...) + ")*"} }
func CharIn(s string) R      { return R{s: "[" + regexp.QuoteMeta(s) + "]"} }
func CharNotIn(s string) R   { return R{s: "[^" + regexp.QuoteMeta(s) + "]"} }
func Char(s string) R        { return R{s: regexp.QuoteMeta(s)} }
func Word() R                { return R{s: `\w+`} }
func WordChar() R            { return R{s: `\w`} }
func WordBoundary() R        { return R{s: `\b`} }
func Digit() R               { return R{s: `\d`} }
func Whitespace() R          { return R{s: `\s`} }
func Tab() R                 { return R{s: `\t`} }
func Linefeed() R            { return R{s: `\n`} }
func CarriageReturn() R      { return R{s: `\r`} }
func Letter() R              { return R{s: `[A-Za-z]`} }
func Lowercase() R           { return R{s: `[a-z]`} }
func Uppercase() R           { return R{s: `[A-Z]`} }

func render(x any) string {
	switch v := x.(type) {
	case R:
		return v.s
	case string:
		return regexp.QuoteMeta(v)
	default:
		panic("invalid input")
	}
}

func concat(xs ...any) string {
	var b strings.Builder
	for _, x := range xs {
		b.WriteString(render(x))
	}
	return b.String()
}

func concatEsc(xs ...any) string {
	var b strings.Builder
	for _, x := range xs {
		switch v := x.(type) {
		case R:
			b.WriteString(v.s)
		case string:
			b.WriteString(regexp.QuoteMeta(v))
		default:
			panic("invalid input")
		}
	}
	return b.String()
}

func join(xs []any, sep string) string {
	var s []string
	for _, x := range xs {
		s = append(s, render(x))
	}
	return strings.Join(s, sep)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
