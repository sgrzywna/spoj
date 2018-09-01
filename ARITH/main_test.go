package main

import (
	"reflect"
	"testing"
)

type TestFR struct {
	first []byte
	res   []byte
}

type TestFSR struct {
	first  []byte
	second []byte
	res    []byte
}

type TestP struct {
	line     []byte
	first    []byte
	second   []byte
	operator byte
}

var AddTests = []TestFSR{
	{[]byte{0}, []byte{0}, []byte{0}},
	{[]byte{1}, []byte{1}, []byte{2}},
	{[]byte{9}, []byte{9}, []byte{1, 8}},
	{[]byte{9, 9}, []byte{9, 9}, []byte{1, 9, 8}},
}

func TestAdd(t *testing.T) {
	for _, tst := range AddTests {
		res := add(tst.first, tst.second)
		if !reflect.DeepEqual(res, tst.res) {
			t.Error(
				"For", tst.first, tst.second,
				"expected", tst.res,
				"got", res,
			)
		}
	}
}

var SubTests = []TestFSR{
	{[]byte{0}, []byte{0}, []byte{0}},
	{[]byte{9}, []byte{9}, []byte{0}},
	{[]byte{9, 9}, []byte{9, 9}, []byte{0, 0}},
	{[]byte{3, 4}, []byte{2, 1}, []byte{1, 3}},
	{[]byte{3, 4}, []byte{2, 6}, []byte{0, 8}},
}

func TestSub(t *testing.T) {
	for _, tst := range SubTests {
		res := sub(tst.first, tst.second)
		if !reflect.DeepEqual(res, tst.res) {
			t.Error(
				"For", tst.first, tst.second,
				"expected", tst.res,
				"got", res,
			)
		}
	}
}

var MulPartialTests = []TestFSR{
	{[]byte{0}, []byte{0}, []byte{0}},
	{[]byte{1}, []byte{2}, []byte{2}},
	{[]byte{2, 3}, []byte{1}, []byte{2, 3}},
	{[]byte{1, 2, 3}, []byte{9}, []byte{1, 1, 0, 7}},
}

func TestMulPartial(t *testing.T) {
	for _, tst := range MulPartialTests {
		res := mulPartial(tst.first, tst.second[0])
		if !reflect.DeepEqual(res, tst.res) {
			t.Error(
				"For", tst.first, tst.second,
				"expected", tst.res,
				"got", res,
			)
		}
	}
}

var MulTests = []TestFSR{
	{[]byte{0}, []byte{0}, []byte{0}},
	{[]byte{1}, []byte{2, 3}, []byte{2, 3}},
	{[]byte{3}, []byte{2, 1}, []byte{6, 3}},
	{[]byte{1, 2}, []byte{3, 4}, []byte{4, 0, 8}},
	{[]byte{1, 2, 3, 4, 5}, []byte{6, 7, 8, 9, 0}, []byte{8, 3, 8, 1, 0, 2, 0, 5, 0}},
}

func TestMul(t *testing.T) {
	for _, tst := range MulTests {
		res, _ := mul(tst.first, tst.second)
		if !reflect.DeepEqual(res, tst.res) {
			t.Error(
				"For", tst.first, tst.second,
				"expected", tst.res,
				"got", res,
			)
		}
	}
}

var ToBytesTests = []TestFR{
	{[]byte{'0'}, []byte{0}},
	{[]byte{'1'}, []byte{1}},
	{[]byte{'1', '2', '3'}, []byte{1, 2, 3}},
}

func TestToBytes(t *testing.T) {
	for _, tst := range ToBytesTests {
		res := toBytes(tst.first)
		if !reflect.DeepEqual(res, tst.res) {
			t.Error(
				"For", string(tst.first),
				"expected", tst.res,
				"got", res,
			)
		}
	}
}

var ToCharsTests = []TestFR{
	{[]byte{'0'}, []byte{0}},
	{[]byte{'1'}, []byte{1}},
	{[]byte{'1', '2', '3'}, []byte{1, 2, 3}},
}

func TestToChars(t *testing.T) {
	for _, tst := range ToCharsTests {
		res := toChars(tst.res)
		if !reflect.DeepEqual(res, tst.first) {
			t.Error(
				"For", string(tst.res),
				"expected", tst.first,
				"got", res,
			)
		}
	}
}

var ParseTests = []TestP{
	{[]byte{'1', '+', '2'}, []byte{1}, []byte{2}, '+'},
	{[]byte{'1', '+', '2', '3'}, []byte{1}, []byte{2, 3}, '+'},
	{[]byte{'1', '2', '+', '3', '4'}, []byte{1, 2}, []byte{3, 4}, '+'},
}

func TestParse(t *testing.T) {
	for _, tst := range ParseTests {
		orgline := make([]byte, len(tst.line))
		reflect.Copy(reflect.ValueOf(orgline), reflect.ValueOf(tst.line))
		first, second, operator := parse(tst.line)
		if !reflect.DeepEqual(first, tst.first) || !reflect.DeepEqual(first, tst.first) || operator != tst.operator {
			t.Error(
				"For", string(orgline),
				"expected", tst.first, tst.operator, tst.second,
				"got", first, operator, second,
			)
		}
	}
}
