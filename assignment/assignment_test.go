package assignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	cases := []struct {
		param1     uint32
		param2     uint32
		sumResult  uint32
		isOverflow bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
	}

	for _, tc := range cases {

		t.Run(
			fmt.Sprintf("%v", tc),
			func(t *testing.T) {
				sum, overflow := AddUint32(tc.param1, tc.param2)
				assert.Equal(t, sum, tc.sumResult)
				assert.Equal(t, tc.isOverflow, overflow)
			})
	}

}

func TestCeilNumber(t *testing.T) {
	cases := []struct {
		param  float64
		result float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}

	for _, tc := range cases {
		t.Run(
			fmt.Sprintf("%v", tc),
			func(t *testing.T) {
				point := CeilNumber(tc.param)
				assert.Equal(t, point, tc.result)
			})
	}

}

func TestAlphabetSoup(t *testing.T) {

	cases := []struct {
		param  string
		result string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}

	for _, tc := range cases {
		t.Run(
			fmt.Sprintf("%v", tc),
			func(t *testing.T) {
				result := AlphabetSoup(tc.param)
				assert.Equal(t, result, tc.result)

			})
	}

}

func TestStringMask(t *testing.T) {

	cases := []struct {
		param  string
		count  uint
		result string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 2, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for _, tc := range cases {
		t.Run(
			fmt.Sprintf("%v", tc),
			func(t *testing.T) {
				result := StringMask(tc.param, tc.count)
				assert.Equal(t, result, tc.result)
			})
	}

}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	cases := []struct {
		words  [2]string
		result string
	}{
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}

	for _, tc := range cases {

		t.Run(
			fmt.Sprintf("%v", tc),
			func(t *testing.T) {
				result := WordSplit(tc.words)
				assert.Equal(t, result, tc.result)
			})
	}

}

func TestVariadicSet(t *testing.T) {

	cases := []struct {
		given    []interface{}
		expected []interface{}
	}{
		{
			[]interface{}{4, 2, 5, 4, 2, 4},
			[]interface{}{4, 2, 5},
		},
		{
			[]interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			[]interface{}{"bootcamp", "rocks!", "really"},
		},
		{
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second"},
		},
	}

	for _, v := range cases {
		t.Run(
			fmt.Sprintf("%v", v),
			func(t *testing.T) {
				result := VariadicSet(v.given)
				assert.Equal(t, v.expected, result)
			})
	}
}
