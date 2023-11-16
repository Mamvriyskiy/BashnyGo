package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func intValue(v int) *int{
	return &v
}

func boolValue(flag bool) *bool {
	return &flag
}

func TestUpdateStrOptions(t *testing.T) {
	var testCase = []struct {
		nameTest string
		strTest string
		opt options
		expected string
	} {
		{"TestIgnoreRegisterA", "abcdf", 
			options{ 
				ignoreRegister : boolValue(true),
				ignoreFirstFields : intValue(0),
				ignoreFirstSymbol : intValue(0),
				}, 
		"abcdf"},
		{"TestIgnoreRegisterB", "Ab CdF", 
			options{ 
				ignoreRegister : boolValue(true),
				ignoreFirstFields : intValue(0),
				ignoreFirstSymbol : intValue(0),
				}, 
		"ab cdf"},
		{"TestIgnoreRegisterC", "ABCDF", 
			options{ 
				ignoreRegister : boolValue(true),
				ignoreFirstFields : intValue(0),
				ignoreFirstSymbol : intValue(0),
				}, 
		"abcdf"},
		{"TestignoreFirstFieldsA", "a b c d f", 
			options{ 
				ignoreRegister : boolValue(false),
				ignoreFirstFields : intValue(1),
				ignoreFirstSymbol : intValue(0),
				}, 
		"b c d f"},
		{"TestignoreFirstFieldsB", "ab c d f", 
			options{ 
				ignoreRegister : boolValue(false),
				ignoreFirstFields : intValue(2),
				ignoreFirstSymbol : intValue(0),
				}, 
		"d f"},
		{"TestignoreFirstFieldsC", "ab cd f", 
			options{ 
				ignoreRegister : boolValue(false),
				ignoreFirstFields : intValue(2),
				ignoreFirstSymbol : intValue(0),
				}, 
		"f"},
		{"TestignoreFirstSymbolA", "abcdf", 
		options{ 
			ignoreRegister : boolValue(false),
			ignoreFirstFields : intValue(0),
			ignoreFirstSymbol : intValue(1),
			}, 
		"bcdf"},
		{"TestignoreFirstSymbolB", "ab cdf", 
		options{ 
			ignoreRegister : boolValue(false),
			ignoreFirstFields : intValue(0),
			ignoreFirstSymbol : intValue(3),
			}, 
		"cdf"},
		{"TestignoreFirstSymbolC", "ab cdf", 
		options{ 
			ignoreRegister : boolValue(false),
			ignoreFirstFields : intValue(0),
			ignoreFirstSymbol : intValue(5),
			}, 
		"f"},
		{"TestA", "ABC df", 
		options{ 
			ignoreRegister : boolValue(true),
			ignoreFirstFields : intValue(1),
			ignoreFirstSymbol : intValue(1),
			}, 
		"f"},
		{"TestB", "ABC df FD", 
		options{ 
			ignoreRegister : boolValue(true),
			ignoreFirstFields : intValue(1),
			ignoreFirstSymbol : intValue(1),
			}, 
		"f fd"},
		{"TestB", "ABC df FD", 
		options{ 
			ignoreRegister : boolValue(true),
			ignoreFirstFields : intValue(2),
			ignoreFirstSymbol : intValue(0),
			}, 
		"fd"},
		}

	for _, tc := range testCase {
		t.Run(tc.nameTest, func(t *testing.T) {
			result := updateStrOptions(tc.strTest, tc.opt)
			assert.Equal(t, tc.expected, result, "string should be equal")
		})
	}
}
