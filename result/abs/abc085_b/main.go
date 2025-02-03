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
	var n int
	n = nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	unique := make(map[int]bool)
	for _, j := range a {
		unique[j] = true
	}
	fmt.Println(len(unique))
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
