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

func handleOperator(c int32, st *IStack, pn *[]string) {
	if (*st).Length() == 0 {
		(*st).Push(c)
	} else {
		h := (*st).Top()

		opPriority := getOpPriority()

		if opPriority[h.(int32)] >= opPriority[c] {
			*pn = append(*pn, string(h.(int32)))
			(*st).Pop()
		}

		(*st).Push(c)
	}
}

func handleRightBracket(st *IStack, pn []string) ([]string, error) {
	for {
		if (*st).Length() == 0 {
			return pn, errors.New("left bracket not found")
		}

		h := (*st).Top()
		(*st).Pop()

		if getBracketInfo(h.(int32)) == lBracket {
			break
		}

		pn = append(pn, string(h.(int32)))
	}

	return pn, nil
}

func handleBracketOrOperator(c int32, num *string, st *IStack, pn []string) ([]string, error) {
	if len(*num) != 0 {
		pn = append(pn, *num)

		*num = ""
	}

	switch {
	case isOperator(c):
		handleOperator(c, st, &pn)
	case getBracketInfo(c) == lBracket:
		(*st).Push(c)
	case getBracketInfo(c) == rBracket:
		var e error
		pn, e = handleRightBracket(st, pn)
		if e != nil {
			return pn, e
		}
	}

	return pn, nil
}

func unwrapStack(num *string, st *IStack, pn []string) []string {
	if len(*num) != 0 {
		pn = append(pn, *num)
	}

	for {
		if (*st).Length() == 0 {
			break
		}

		h := (*st).Top()

		pn = append(pn, string(h.(int32)))

		(*st).Pop()
	}

	return pn
}

func evaluateWithOperator(c int32, arg1 float64, arg2 float64, st *IStack) error {
	switch c {
	case '+':
		(*st).Push(arg1 + arg2)
	case '-':
		(*st).Push(arg1 - arg2)
	case '/':
		(*st).Push(arg1 / arg2)
	case '*':
		(*st).Push(arg1 * arg2)
	default:
		return errors.New("unknown operator")
	}

	return nil
}

func extractArguments(st *IStack) (arg1 float64, arg2 float64, e error) {
	arg2, e = getStackValue(st)
	if e != nil {
		return
	}

	arg1, e = getStackValue(st)
	if e != nil {
		return
	}

	return
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
			var e error
			pn, e = handleBracketOrOperator(c, &num, &stack, pn)
			if e != nil {
				return pn, e
			}
		case c == '\n':
			pn = unwrapStack(&num, &stack, pn)
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
			arg1, arg2, e := extractArguments(&stack)
			if e != nil {
				return result, e
			}
			e = evaluateWithOperator(c, arg1, arg2, &stack)
			if e != nil {
				return result, e
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