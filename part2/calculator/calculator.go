package calculator

import (
	"errors"
	"strconv"
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

					if opPriority[h.(int32)] >= opPriority[c] {
						pn = append(pn, string(h.(int32)))
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

					if getBracketInfo(h.(int32)) == lBracket {
						break
					}

					pn = append(pn, string(h.(int32)))
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

				pn = append(pn, string(h.(int32)))

				stack.Pop()
			}

		default:
			return pn, errors.New("invalid character")
		}
	}

	return pn, nil
}

func getStackValue(st *IStack) (v float64, e error) {
	if (*st).Length() == 0 {
		return 0, errors.New("invalid order of operations")
	}

	v = (*st).Top().(float64)
	(*st).Pop()

	return
}

func Calculate(s string) (float64, error) {
	var result float64 = 0

	pn, e := getPolishNotation(s)

	if e != nil {
		return result, errors.New("polish notation conversion failure")
	}

	stack := NewStack()

	for _, v := range pn {
		c := int32(v[0])
		switch {
		case isNumber(c):
			num, e := strconv.Atoi(v)

			if e != nil {
				return result, errors.New("invalid number")
			}

			stack.Push(float64(num))
		case isOperator(c):
			arg2, e := getStackValue(&stack)

			if e != nil {
				return result, errors.New("invalid order of operations")
			}

			arg1, e := getStackValue(&stack)

			if e != nil {
				return result, errors.New("invalid order of operations")
			}

			switch c {
			case '+':
				stack.Push(arg1 + arg2)
			case '-':
				stack.Push(arg1 - arg2)
			case '/':
				stack.Push(arg1 / arg2)
			case '*':
				stack.Push(arg1 * arg2)
			}
		default:
			return result, errors.New("unknown character occurred")
		}
	}

	if stack.Length() == 0 {
		return result, errors.New("result not found")
	}

	result = stack.Top().(float64)

	return result, nil
}