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
	sc.Scan()
	nums := make([]int, n)
	for i, s := range strings.Split(sc.Text(), " ") {
		nums[i], _ = strconv.Atoi(s)
	}

	count := 0
	for {
		// 全ての数が偶数かチェック
		allEven := true
		for _, num := range nums {
			if num % 2 != 0 {
				allEven = false
				break
			}
		}

		if !allEven {
			break
		}

		// 全ての数を2で割る
		for i := range nums {
			nums[i] = nums[i] / 2
		}
		count++
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
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, err := os.ReadFile("./input")
		if err != nil {
			panic(err)
		}
		sc = bufio.NewScanner(strings.NewReader(string(b)))
	}
}
