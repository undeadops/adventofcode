package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// Load day1.input
	file, err := os.Open("./day2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var depth int
	var horiz int

	// Set Last initially to 0
	depth = 0
	horiz = 0

	// Run through each line of input
	for scanner.Scan() {
		w := strings.Fields(scanner.Text())
		action := w[0]
		value, _ := strconv.Atoi(w[1])

		switch action {
		case "down":
			depth = depth + value
		case "up":
			depth = depth - value
		case "forward":
			horiz = horiz + value
		}
	}

	fmt.Println("Depth: ", depth)
	fmt.Println("Horizontal: ", horiz)
	fmt.Printf("Combined: %d\n", depth*horiz)
}

func part2() {
	file, err := os.Open("./day2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var depth int
	var horiz int
	var aim int

	// Set everything to 0
	depth = 0
	horiz = 0
	aim = 0

	for scanner.Scan() {
		w := strings.Fields(scanner.Text())
		action := w[0]
		value, _ := strconv.Atoi(w[1])

		switch action {
		case "down":
			aim = aim + value
		case "up":
			aim = aim - value
		case "forward":
			horiz = horiz + value
			depth = depth + (aim * value)
		}
	}

	fmt.Println("Depth: ", depth)
	fmt.Println("Horizontal: ", horiz)
	fmt.Println("AIM: ", aim)
	fmt.Printf("Combined: %d\n", depth*horiz)
}

func main() {
	part1()
	part2()
}
