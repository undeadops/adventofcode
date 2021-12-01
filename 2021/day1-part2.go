package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Individual lines
	var inputs []int

	file, err := os.Open("./day1.input")
	if err != nil {
		panic("Unable to open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, v)
	}

	last := 0
	count := 0
	for x, _ := range inputs {
		t := inputs[x : x+3]
		s := t[0] + t[1] + t[2]

		// fix off by one error
		if last == 0 {
			last = s
		}

		if !(t[1] == 0 || t[2] == 0) {
			if s > last {
				count++
				fmt.Printf("%d: increased \n", s)
			} else {
				fmt.Printf("%d: decreased \n", s)
			}
		}
		last = s
	}

	fmt.Printf("Count: %d\n", count)
}
