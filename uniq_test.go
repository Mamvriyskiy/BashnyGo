package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStrOptions(t *testing.T) {
	var testCase = []struct {
		nameTest string
		strTest string
		opt options
		expected string
	} {
		{
			"TestIgnoreRegisterA", "abcdf", 
			options{ 
				ignoreRegister : true,
				ignoreFirstFields : 0,
				ignoreFirstSymbol : 0,
			}, 
			"abcdf"},
		{
			"TestIgnoreRegisterB", "Ab CdF", 
			options{ 
				ignoreRegister : true,
				ignoreFirstFields : 0,
				ignoreFirstSymbol : 0,
			}, 
			"ab cdf"},
		{
			"TestIgnoreRegisterC", "ABCDF", 
			options{ 
				ignoreRegister : true,
				ignoreFirstFields : 0,
				ignoreFirstSymbol : 0,
			}, 
			"abcdf"},
		{
			"TestignoreFirstFieldsA", "a b c d f", 
			options{ 
				ignoreRegister : false,
				ignoreFirstFields : 1,
				ignoreFirstSymbol : 0,
			}, 
			"b c d f"},
		{
			"TestignoreFirstFieldsB", "ab c d f", 
			options{ 
				ignoreRegister : false,
				ignoreFirstFields : 2,
				ignoreFirstSymbol : 0,
			}, 
			"d f"},
		{
			"TestignoreFirstFieldsC", "ab cd f", 
			options{ 
				ignoreRegister : false,
				ignoreFirstFields : 2,
				ignoreFirstSymbol : 0,
			}, 
			"f"},
		{
			"TestignoreFirstSymbolA", "abcdf", 
			options{ 
				ignoreRegister : false,
				ignoreFirstFields : 0,
				ignoreFirstSymbol : 1,
			}, 
			"bcdf"},
		{
			"TestignoreFirstSymbolB", "ab cdf", 
			options{ 
				ignoreRegister : false,
				ignoreFirstFields : 0,
				ignoreFirstSymbol : 3,
			}, 
			"cdf"},
		{
			"TestignoreFirstSymbolC", "ab cdf", 
			options{ 
				ignoreRegister : false,
				ignoreFirstFields : 0,
				ignoreFirstSymbol : 5,
			}, 
			"f"},
		{
			"TestA", "ABC df", 
			options{ 
				ignoreRegister : true,
				ignoreFirstFields : 1,
				ignoreFirstSymbol : 1,
			}, 
			"f"},
		{
			"TestB", "ABC df FD", 
			options{ 
				ignoreRegister : true,
				ignoreFirstFields : 1,
				ignoreFirstSymbol : 1,
				}, 
			"f fd"},
		{
			"TestB", "ABC df FD", 
			options{ 
				ignoreRegister : true,
				ignoreFirstFields : 2,
				ignoreFirstSymbol : 0,
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
