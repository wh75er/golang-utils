package iop

import (
	"bufio"
	"os"
	"uniq/args"
)

func ReadInput(opts args.IoOptions) ([]string, error){
	var data []string

	inStream := os.Stdin

	if opts.InFilename != "" {
		f, e := os.OpenFile(opts.InFilename, os.O_RDONLY, os.FileMode(0644))

		if e != nil {
			return data, e
		}

		defer f.Close()


		inStream = f
	}

	scanner := bufio.NewScanner(inStream)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if e := scanner.Err(); e != nil {
		return data, e
	}

	return data, nil
}