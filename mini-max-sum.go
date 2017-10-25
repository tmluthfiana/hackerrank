package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Mini-Max Sum
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	var in []int
	for _, a := range aa {
		num, _ := strconv.Atoi(a)
		in = append(in, num)
	}

	var a = make([]int, 5)
	var sum int
	sum = 0

	for i := 0; i < len(a); i++ {
		a[i] = in[i]
		sum += a[i]
	}

	sort.Ints(a)
	fmt.Println((sum - a[len(a)-1]), (sum - a[0]))

}
