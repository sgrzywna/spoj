package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1111111)
	scanner.Buffer(buf, len(buf))
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		input := scanner.Text()
		number := inc(normalize([]byte(input)))
		left, right := findPalindrome(number)
		os.Stdout.Write(left)
		os.Stdout.Write(right)
		os.Stdout.Write([]byte("\n"))
	}
}

func normalize(number []byte) []byte {
	h := 0
	t := len(number) - 1
	// skip tail spaces
	for t >= 0 && number[t] == ' ' {
		t--
	}
	// skip head zeroes and spaces
	for h <= t && (number[h] == '0' || number[h] == ' ') {
		h++
	}
	if h > t {
		return []byte("0")
	}
	return number[h : t+1]
}

func inc(number []byte) []byte {
	t := len(number) - 1
	for t >= 0 {
		if number[t] < '9' {
			number[t]++
			return number
		}
		number[t] = '0'
		t--
	}
	newnumber := make([]byte, len(number)+1)
	copy(newnumber[1:], number)
	newnumber[0] = '1'
	return newnumber
}

func findPalindrome(number []byte) (left, right []byte) {
	h := 0
	t := len(number) - 1
	len := t - h
	// single digit case
	if len == 0 {
		return number[h : h+1], []byte{}
	}
	// millions of digits
	half := len / 2
	even := len % 2
	left = number[h : h+half+1]
	right = number[h+half+even : t+1]
	rleft := reverse(left)
	cmp := compare(right, rleft)
	if cmp == 0 {
		right = right[1-even:]
	} else if cmp == -1 {
		right = rleft[1-even:]
	} else {
		left = inc(left)
		rleft = reverse(left)
		right = rleft[1-even:]
	}
	return left, right
}

// compare returns 1 when left > right, -1 when left < right, 0 for equal
func compare(left, right []byte) int {
	h := 0
	t := len(left) - 1
	for h <= t {
		if left[h] > right[h] {
			return 1
		}
		if left[h] < right[h] {
			return -1
		}
		h++
	}
	return 0
}

// reverse returns reversed copy of input
func reverse(number []byte) []byte {
	res := make([]byte, len(number))
	h := 0
	t := len(number) - 1
	for t >= 0 {
		res[h] = number[t]
		h++
		t--
	}
	return res
}
