package args

import "errors"

type Options struct {
	SkipFields int
	SkipChars int
	Count bool
	Repeated bool
	Unique bool
	IgnoreCase bool
}

func (o Options) IsValid() error {
	var conflictCount byte
	for _, v := range []bool{o.Repeated, o.Unique, o.Count} {
		if v {
			conflictCount++

			if conflictCount > 1 {
				return errors.New("conflicting options")
			}
		}
	}

	return nil
}

type OptionBuilder interface {
	SkipFields(int) OptionBuilder
	SkipChars(int) OptionBuilder
	Count() OptionBuilder
	Repeated() OptionBuilder
	Unique() OptionBuilder
	IgnoreCase() OptionBuilder
	Build() Options
}

func (o *optionBuilder) SkipFields(n int) OptionBuilder {
	o.skipFields = n
	return o
}

func (o *optionBuilder) SkipChars(n int) OptionBuilder {
	o.skipChars = n
	return o
}

func (o *optionBuilder) Count() OptionBuilder {
	o.count = true
	return o
}

func (o *optionBuilder) Repeated() OptionBuilder {
	o.repeated = true
	return o
}

func (o *optionBuilder) Unique() OptionBuilder {
	o.unique = true
	return o
}

func (o *optionBuilder) IgnoreCase() OptionBuilder {
	o.ignoreCase = true
	return o
}

func (o *optionBuilder) Build() Options {
	return Options{
		o.skipFields,
		o.skipChars,
		o.count,
		o.repeated,
		o.unique,
		o.ignoreCase,
	}
}

type optionBuilder struct {
	skipFields int
	skipChars int
	count bool
	repeated bool
	unique bool
	ignoreCase bool
}

func NewOptions() OptionBuilder {
	return &optionBuilder{}
}