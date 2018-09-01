package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		input := scanner.Text()
		first, second, operator := parse([]byte(input))
		var res []byte
		switch operator {
		case '+':
			res = add(first, second)
			render(first, second, operator, res)
		case '-':
			res = sub(first, second)
			render(first, second, operator, res)
		case '*':
			res, partials := mul(first, second)
			renderMul(first, second, res, partials)
		}
		fmt.Fprintln(os.Stdout)
	}
}

func render(first, second []byte, operator byte, res []byte) {
	first = toChars(first)
	second = append([]byte{operator}, toChars(second)...)
	res = toChars(trimLeadingZeroes(res))
	length := len(first)
	if len(second) > length {
		length = len(second)
	}
	if len(res) > length {
		length = len(res)
	}
	fmt.Fprintln(os.Stdout, string(renderWithIndent(first, length)))
	fmt.Fprintln(os.Stdout, string(renderWithIndent(second, length)))
	fmt.Fprintln(os.Stdout, string(renderWithIndent(renderDashes(maxInt(len(second), len(res))), length)))
	fmt.Fprintln(os.Stdout, string(renderWithIndent(res, length)))
}

func renderMul(first, second, res []byte, partials [][]byte) {
	first = toChars(first)
	second = append([]byte{'*'}, toChars(second)...)
	res = toChars(trimLeadingZeroes(res))
	length := len(first)
	if len(second) > length {
		length = len(second)
	}
	if len(res) > length {
		length = len(res)
	}

	fmt.Fprintln(os.Stdout, string(renderWithIndent(first, length)))
	fmt.Fprintln(os.Stdout, string(renderWithIndent(second, length)))

	if len(partials) > 1 {
		partials[0] = trimLeadingZeroes(partials[0])
		dashesLength := maxInt(len(second), len(partials[0]))
		fmt.Fprintln(os.Stdout, string(renderWithIndent(renderDashes(dashesLength), length)))
		for i, partial := range partials {
			fmt.Fprintln(os.Stdout, string(renderWithIndent(toChars(trimLeadingZeroes(partial)), length-i)))
		}
	}

	fmt.Fprintln(os.Stdout, string(renderWithIndent(renderDashes(len(res)), length)))
	fmt.Fprintln(os.Stdout, string(renderWithIndent(res, length)))
}

func renderWithIndent(data []byte, length int) []byte {
	res := make([]byte, length)
	shift := length - len(data)
	copy(res[shift:], data)
	for i := 0; i < shift; i++ {
		res[i] = ' '
	}
	return res
}

func trimLeadingZeroes(number []byte) []byte {
	h := 0
	t := len(number) - 1
	for number[h] == 0 && h < t {
		h++
	}
	return number[h:]
}

func renderDashes(length int) []byte {
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = '-'
	}
	return res
}

func parse(line []byte) ([]byte, []byte, byte) {
	length := len(line)
	op := 0
	for h := 0; op == 0 && h < length; h++ {
		if line[h] == '+' || line[h] == '-' || line[h] == '*' {
			op = h
		}
	}
	return toBytes(line[:op]), toBytes(line[op+1:]), line[op]
}

func toBytes(number []byte) []byte {
	for h := 0; h < len(number); h++ {
		number[h] -= '0'
	}
	return number
}

func toChars(number []byte) []byte {
	for h := 0; h < len(number); h++ {
		number[h] += '0'
	}
	return number
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func normalizeSlice(number []byte, length int) []byte {
	if len(number) == length {
		return number
	}
	res := make([]byte, length)
	shift := length - len(number)
	copy(res[shift:], number)
	return res
}

func add(first, second []byte) []byte {
	maxDigits := maxInt(len(first), len(second))
	res := make([]byte, maxDigits)
	first = normalizeSlice(first, maxDigits)
	second = normalizeSlice(second, maxDigits)
	carry := byte(0)
	for t := maxDigits - 1; t >= 0; t-- {
		sum := first[t] + second[t] + carry
		carry = sum / 10
		res[t] = sum % 10
	}
	if carry > 0 {
		newres := make([]byte, len(res)+1)
		copy(newres[1:], res)
		newres[0] = carry
		res = newres
	}
	return res
}

func sub(first, second []byte) []byte {
	maxDigits := maxInt(len(first), len(second))
	res := make([]byte, maxDigits)
	first = normalizeSlice(first, maxDigits)
	second = normalizeSlice(second, maxDigits)
	carry := 0
	for t := maxDigits - 1; t >= 0; t-- {
		sum := int(first[t]) - int(second[t]) - carry
		if sum < 0 {
			carry = 1
			res[t] = byte(sum + 10)
		} else {
			carry = 0
			res[t] = byte(sum)
		}
	}
	return res
}

func mul(first, second []byte) ([]byte, [][]byte) {
	length := len(second)
	partials := make([][]byte, length)
	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		partials[i] = mulPartial(first, second[j])
	}
	res := make([]byte, len(partials[0]))
	for i := 0; i < len(partials); i++ {
		partial := make([]byte, len(partials[i])+i)
		copy(partial, partials[i])
		res = add(res, partial)
	}
	return res, partials
}

func mulPartial(first []byte, second byte) []byte {
	length := len(first)
	res := make([]byte, length)
	carry := byte(0)
	for t := length - 1; t >= 0; t-- {
		m := first[t]*second + carry
		res[t] = m % 10
		carry = m / 10
	}
	if carry > 0 {
		newres := make([]byte, len(res)+1)
		copy(newres[1:], res)
		newres[0] = carry
		res = newres
	}
	return res
}
