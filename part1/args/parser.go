package args

import (
	"errors"
	"flag"
)

type IoOptions struct {
	InFilename  string
	OutFilename string
}

func ParseArgs() (Options, IoOptions, error) {
	var e error = nil

	var (
		count = flag.Bool("c", false, "prefix lines by the number of occurrences")
		repeated = flag.Bool("d", false, "only print duplicate lines, one for each group")
		unique = flag.Bool("u", false, "only print unique lines")
		skipFields = flag.Int("f", 0, "avoid comparing the first N fields")
		skipChars = flag.Int("s", 0, "avoid comparing the first N characters")
		ignoreCase = flag.Bool("i", false, "ignore differences in case when comparing")
	)

	flag.Parse()

	builder := NewOptions()

	if *count {
		builder.Count()
	}
	if *repeated {
		builder.Repeated()
	}
	if *unique {
		builder.Unique()
	}
	if *ignoreCase {
		builder.IgnoreCase()
	}
	builder.SkipFields(*skipFields)
	builder.SkipChars(*skipChars)

	opts := builder.Build()

	iopts := IoOptions{}

	args := flag.Args()

	switch len(args) {
	case 0:
		break
	case 1:
		iopts.InFilename = args[0]
	case 2:
		iopts.InFilename = args[0]
		iopts.OutFilename = args[1]
	default:
		e = errors.New("too many arguments")
	}

	return opts, iopts, e
}
