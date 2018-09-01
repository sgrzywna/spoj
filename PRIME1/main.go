package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cases, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	for i := 0; i < cases; i++ {
		scanner.Scan()
		mn := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		m, _ := strconv.Atoi(mn[0])
		n, _ := strconv.Atoi(mn[1])
		for j := m; j <= n; j++ {
			if isPrime(j) {
				fmt.Fprintln(os.Stdout, j)
			}
		}
		fmt.Fprintln(os.Stdout)
	}
}
