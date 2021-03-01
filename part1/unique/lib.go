package unique

import (
	"errors"
	"strconv"
	"strings"
	"uniq/args"
)

const (
	Default = iota
	countMode
	repeatedMode
	uniqueMode
)

type line struct {
	msg string
	count int
}

func detectMode(opts args.Options) int {
	mode := Default

	switch {
	case opts.Count:
		mode = countMode
	case opts.Repeated:
		mode = repeatedMode
	case opts.Unique:
		mode = uniqueMode
	}

	return mode
}

func Map(f func(line) string, l []line) []string {
	lm := make([]string, len(l))
	for i, v := range l {
		lm[i] = f(v)
	}
	return lm
}

func Filter(f func(line) bool, s []line) []string {
	lm := make([]string, 0)
	for _, v := range s {
		if f(v) {
			lm = append(lm, v.msg)
		}
	}

	return lm
}

func merge(l []line, m int) ([]string, error) {
	if m != uniqueMode && m != repeatedMode {
		return nil, errors.New("unknown merge mode")
	}

	var lm []string

	switch m {
	case uniqueMode:
		lm = Filter(func (v line) bool {
			if v.count == 0 {
				return true
			}
			return false
		}, l)
	case repeatedMode:
		lm = Filter(func (v line) bool {
			if v.count != 0 {
				return true
			}
			return false
		}, l)
	}

	return lm, nil
}

func modifyString(s string, fields int, chars int, ignoreCase bool) string {
	if ignoreCase {
		s = strings.ToLower(s)
	}

	words := strings.Split(s, " ")

	if fields > len(words) {
		fields = len(words)
	}

	words = words[fields:]

	modifiedString := strings.Join(words, " ")

	if chars > len(modifiedString) {
		chars = len(modifiedString)
	}

	modifiedString = modifiedString[chars:]

	return modifiedString
}

func Uniqualize(data []string, opts args.Options) ([]string, error) {
	var result []string

	if len(data) == 0 {
		return result, errors.New("input is empty")
	}

	if e := opts.IsValid(); e != nil {
		return result, e
	}

	var lines []line

	lines = append(lines, line{data[0], 0})
	lastLine := modifyString(data[0], opts.SkipFields, opts.SkipChars, opts.IgnoreCase)

	for _, v := range data[1:] {
		currLine := modifyString(v, opts.SkipFields, opts.SkipChars, opts.IgnoreCase)
		if lastLine == currLine {
			lines[len(lines) - 1].count++
		} else {
			lines = append(lines, line{v, 0})
			lastLine = currLine
		}
	}

	mode := detectMode(opts)

	switch mode {
	case countMode:
		result = Map(func (l line) string {
			return strings.Join([]string{strconv.Itoa(l.count + 1), l.msg}, " ")
		}, lines)
	default:
		var e error = nil
		result, e = merge(lines, mode)
		if e != nil {
			return result, e
		}
	}

	return result, nil
}