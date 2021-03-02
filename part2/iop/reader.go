package iop

import (
	"bufio"
	"io"
	"errors"
)

func ReadInput(in io.Reader) (string, error){
	r := bufio.NewReader(in)

	s, e := r.ReadString('\n')

	if e != nil {
		return s, errors.New("delimiter not found")
	}

	return s, nil
}