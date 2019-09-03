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
<<<<<<< HEAD
	for _, a := range aa {
=======
	aaa := strings.Split(lines[0], " ")
	for _, a := range aaa {
>>>>>>> 056adb0a7913cb8bdb5abbae4a24f775999c5e82
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
