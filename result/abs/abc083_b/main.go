package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n := nextInt()
	a := nextInt()
	b := nextInt()

	var total int
	for i := 1; i <= n; i++ {
		sum := sumDigit(i)

		if sum >= a && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}

func sumDigit(x int) int {
	sum := 0

	// 0以下になるまで
	for x > 0 {
		sum += x % 10
		x /= 10 // 少数点以下は切り捨て
	}

	return sum
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, err := os.ReadFile("./input")
		if err != nil {
			panic(err)
		}
		sc = bufio.NewScanner(strings.NewReader(string(b)))
		sc.Split(bufio.ScanWords)
	}
}
