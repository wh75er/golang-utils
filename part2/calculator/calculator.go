package calculator

import (
	"errors"
	"fmt"
	"strings"
)

const (
	lBracket = iota
	rBracket
	notBracket
)

func getOpPriority() map[int32]byte {
	return map[int32]byte {
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
}

func isOperator(c int32) bool {
	if _, ok := getOpPriority()[c]; !ok {
		return false
	}

	return true
}

func isNumber(c int32) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func getBracketInfo(c int32) byte{
	switch c {
	case '(':
		return lBracket
	case ')':
		return rBracket
	}

	return notBracket
}

func getPolishNotation(s string) ([]string, error) {
	pn := make([]string, 0)

	if len(s) == 0 {
		return pn, nil
	}

	s = strings.ReplaceAll(s, " ", "")

	stack := NewStack()

	num := ""

	for _, c := range s {
		switch {
		case isNumber(c):
			num += string(c)
		case getBracketInfo(c) != notBracket || isOperator(c):
			if len(num) != 0 {
				pn = append(pn, num)

				num = ""
			}

			switch {
			case isOperator(c):
				if stack.Length() == 0 {
					stack.Push(c)
				} else {
					h := stack.Top()

					opPriority := getOpPriority()

					if opPriority[h] >= opPriority[c] {
						pn = append(pn, string(h))
						stack.Pop()
					}

					stack.Push(c)
				}
			case getBracketInfo(c) == lBracket:
				stack.Push(c)
			case getBracketInfo(c) == rBracket:
				for {
					if stack.Length() == 0 {
						return pn, errors.New("left bracket not found")
					}

					h := stack.Top()
					stack.Pop()

					if getBracketInfo(h) == lBracket {
						break
					}

					pn = append(pn, string(h))
				}
			}
		case c == '\n':
			if len(num) != 0 {
				pn = append(pn, num)
			}

			for {
				if stack.Length() == 0 {
					break
				}

				h := stack.Top()

				pn = append(pn, string(h))

				stack.Pop()
			}

		default:
			return pn, errors.New("invalid character")
		}
	}

	return pn, nil
}

func Calculate(s string) (float64, error) {
	pn, e := getPolishNotation(s)

	if e != nil {
		return 0, e
	}

	fmt.Println("Polish notation: ", pn)

	return 0, nil
}