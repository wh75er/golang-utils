package unique

import (
	"reflect"
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
		"f",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}
