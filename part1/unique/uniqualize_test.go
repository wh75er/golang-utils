package unique

import (
	"reflect"
	"testing"
	"uniq/args"
)

func TestDefaultBehaviour(t *testing.T) {
	opts := args.NewOptions().Build()

	dataInput := []string{
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

	expected := []string{
		"I love music.",
		"",
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

func TestCountBehaviour(t *testing.T) {
	opts := args.NewOptions().Count().Build()

	dataInput := []string{
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

	expected := []string{
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}

func TestRepeatedBehaviour(t *testing.T) {
	opts := args.NewOptions().Repeated().Build()

	dataInput := []string{
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

	expected := []string{
		"I love music.",
		"I love music of Kartik.",
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

func TestUniqueBehaviour(t *testing.T) {
	opts := args.NewOptions().Unique().Build()

	dataInput := []string{
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

	expected := []string{
		"",
		"Thanks.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}

func TestLowerCaseBehaviour(t *testing.T) {
	opts := args.NewOptions().IgnoreCase().Build()

	dataInput := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik.",
	}

	expected := []string{
		"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}

func TestSkipFieldsBehaviour(t *testing.T) {
	opts := args.NewOptions().SkipFields(1).Build()

	dataInput := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	expected := []string{
		"We love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}

func TestSkipCharsBehaviour(t *testing.T) {
	opts := args.NewOptions().SkipChars(1).Build()

	dataInput := []string{
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	expected := []string{
		"I love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}

func TestConflictingKeysAllFailure(t *testing.T) {
	opts := args.NewOptions().Count().Repeated().Unique().Build()

	dataInput := []string{
		"foo",
	}

	_, e := Uniqualize(dataInput, opts)

	if e == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestConflictingKeysCountRepeatedFailure(t *testing.T) {
	opts := args.NewOptions().Count().Repeated().Build()

	dataInput := []string{
		"foo",
	}

	_, e := Uniqualize(dataInput, opts)

	if e == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestConflictingKeysCountUniqueFailure(t *testing.T) {
	opts := args.NewOptions().Count().Unique().Build()

	dataInput := []string{
		"foo",
	}

	_, e := Uniqualize(dataInput, opts)

	if e == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestConflictingKeysRepeatedUniqueFailure(t *testing.T) {
	opts := args.NewOptions().Repeated().Unique().Build()

	dataInput := []string{
		"foo",
	}

	_, e := Uniqualize(dataInput, opts)

	if e == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestEmptyInputDataBehaviour(t *testing.T) {
	opts := args.NewOptions().SkipChars(1).Build()

	var dataInput []string

	var expected []string

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}

func TestSkipFieldsSkipCharsBehaviour(t *testing.T) {
	opts := args.NewOptions().SkipFields(2).SkipChars(1).Build()

	dataInput := []string{
		"I love music.",
		"A love gusic.",
		"C love Tusic.",
		"",
		"I love husic of Kartik.",
		"We love Xusic of Kartik.",
		"Thanks.",
	}

	expected := []string{
		"I love music.",
		"",
		"I love husic of Kartik.",
		"Thanks.",
	}

	output, e := Uniqualize(dataInput, opts)

	if e != nil {
		t.Fatalf("Unitity failed with error %s", e)
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Check failed. Got: %s\nexpected:%s", output, expected)
	}
}
