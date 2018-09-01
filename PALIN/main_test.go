package main

import (
	"reflect"
	"testing"
)

type TestArrayArray struct {
	first []byte
	res   []byte
}

type TestArrayArrayInt struct {
	first  []byte
	second []byte
	res    int
}

type TestArrayBool struct {
	first []byte
	res   bool
}

var ReverseTests = []TestArrayArray{
	{[]byte{0, 1, 2, 3}, []byte{3, 2, 1, 0}},
	{[]byte{1, 2, 3}, []byte{3, 2, 1}},
}

func TestReverse(t *testing.T) {
	for _, pair := range ReverseTests {
		res := reverse(pair.first)
		if !reflect.DeepEqual(res, pair.res) {
			t.Error(
				"For", pair.first,
				"expected", pair.res,
				"got", res,
			)
		}
	}
}

var IsSmallerTests = []TestArrayArrayInt{
	{[]byte("0"), []byte("1"), -1},
	{[]byte("1"), []byte("0"), 1},
	{[]byte("09"), []byte("10"), -1},
	{[]byte("10"), []byte("09"), 1},
	{[]byte("10"), []byte("10"), 0},
}

func TestIsSmaller(t *testing.T) {
	for _, pair := range IsSmallerTests {
		res := compare(pair.first, pair.second)
		if res != pair.res {
			t.Error(
				"For", pair.first, pair.second,
				"expected", pair.res,
				"got", res,
			)
		}
	}
}

var IncTests = []TestArrayArray{
	{[]byte("0"), []byte("1")},
	{[]byte("1"), []byte("2")},
	{[]byte("9"), []byte("10")},
	{[]byte("10"), []byte("11")},
	{[]byte("999"), []byte("1000")},
}

func TestInc(t *testing.T) {
	for _, pair := range IncTests {
		res := inc(pair.first)
		if !reflect.DeepEqual(res, pair.res) {
			t.Error(
				"For", pair.first,
				"expected", pair.res,
				"got", res,
			)
		}
	}
}

var FindPalindromeTests = []TestArrayArray{
	{[]byte("0"), []byte("0")},
	{[]byte("1"), []byte("1")},
	{[]byte("9"), []byte("9")},
	{[]byte("11"), []byte("11")},
	{[]byte("99"), []byte("99")},
	{[]byte("808"), []byte("808")},
	{[]byte("1111"), []byte("1111")},
	{[]byte("1112"), []byte("1221")},
	{[]byte("2133"), []byte("2222")},
	{[]byte("111111"), []byte("111111")},
	{[]byte("421999767"), []byte("422000224")},
	{[]byte("94187978322"), []byte("94188088149")},
	{[]byte("2133000"), []byte("2133312")},
	{[]byte("9418797832294949494"), []byte("9418797832387978149")},
	{[]byte("10000000000000000000"), []byte("10000000000000000001")},
}

func TestFindPalindrome(t *testing.T) {
	for _, pair := range FindPalindromeTests {
		org := make([]byte, len(pair.first))
		reflect.Copy(reflect.ValueOf(org), reflect.ValueOf(pair.first))
		left, right := findPalindrome(pair.first)
		res := append(left, right...)
		if !reflect.DeepEqual(res, pair.res) {
			t.Error(
				"For", string(org),
				"expected", string(pair.res),
				"got", string(res),
			)
		}
	}
}

var NormalizeTests = []TestArrayArray{
	{[]byte("0"), []byte("0")},
	{[]byte("1"), []byte("1")},
	{[]byte("10"), []byte("10")},
	{[]byte("000"), []byte("0")},
	{[]byte("001"), []byte("1")},
	{[]byte("010"), []byte("10")},
}

func TestNormalize(t *testing.T) {
	for _, pair := range NormalizeTests {
		res := normalize(pair.first)
		if !reflect.DeepEqual(res, pair.res) {
			t.Error(
				"For", pair.first,
				"expected", pair.res,
				"got", res,
			)
		}
	}
}
