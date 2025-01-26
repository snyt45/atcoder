package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
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
