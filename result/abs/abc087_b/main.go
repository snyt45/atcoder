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
	a := nextInt()
	b := nextInt()
	c := nextInt()
	x := nextInt()

	var count, total int

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for h := 0; h <= c; h++ {
				total = 500*i + 100*j + 50*h
				if total == x {
					count++
				}
			}
		}
	}

	fmt.Println(count)
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
	}
}
