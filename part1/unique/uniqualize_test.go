package unique

import (
	"testing"
	"uniq/args"
)

func testStandardBehaviour(t *testing.T) {
	opts := args.NewOptions().Build()

	dataInput := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	expected := []string {
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
	}

	output, e := Uniqualize(dataInput, opts)
}