package goconveycmp

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

const (
	success                = ""
	needExactValues        = "This assertion requires exactly %d comparison values (you provided %d)."
	needNonEmptyCollection = "This assertion requires at least 1 comparison value (you provided 0)."
	needFewerValues        = "This assertion allows %d or fewer comparison values (you provided %d)."
)

// ShouldCmp receives exactly two parameters and ensures that the first is equal to the second with go-cmp.
func ShouldCmp(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}
	diff := cmp.Diff(actual, expected[0])
	if diff == "" {
		return success
	}
	return "Expected no difference, got:\n" + diff
}

// ShouldCmpProto is same  as ShouldCmp but for protocol buffer messages.
func ShouldCmpProto(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}
	diff := cmp.Diff(actual, expected[0], protocmp.Transform())
	if diff == "" {
		return success
	}
	return "Expected no difference, got:\n" + diff
}

func need(needed int, expected []interface{}) string {
	if len(expected) != needed {
		return fmt.Sprintf(needExactValues, needed, len(expected))
	}
	return success
}
