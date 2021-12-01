package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Load day1.input
	file, err := os.Open("./day1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int
	var last int

	// Set Last initially to 0
	last = 0

	// Run through each line of input
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())

		if last == 0 {
			last = v
		}

		if v > last {
			fmt.Printf("%d (increase)\n", v)
			count++
		} else {
			fmt.Printf("%d (decrease)\n", v)
		}

		// Update last
		last = v
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The Count is: %d", count)
}
