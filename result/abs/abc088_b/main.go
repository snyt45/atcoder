package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// スライスを降順に整列させる
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	var alice, bob int

	for i, num := range nums {
		// 偶数のときは、alice（先行）
		if i%2 == 0 {
			alice += num
			// 奇数のときは、bob（後攻）
		} else {
			bob += num
		}
	}

	fmt.Println(alice - bob)
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
