package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Birthday Cake Candles
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	var in []int64
	n, _ := strconv.Atoi(lines[0])
	aaa := strings.Split(lines[1], " ")
	for _, a := range aaa {
		num, _ := strconv.ParseInt(a, 0, 64)
		in = append(in, num)
	}

	var tallest int64
	frequency := 0

	for i := 0; i < n; i++ {
		height := in[i]

		if height > tallest {
			tallest = height
			frequency = 1
		} else if height == tallest {
			frequency++
		}

	}

	fmt.Println(frequency)
}
