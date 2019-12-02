package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputFile = flag.String("inputFile", "inputs/day01.txt", "Relative Path for input file")

func calcfuel(mass int) int {
	var fuel int
	fuel = (int(mass) / int(3)) - int(2)
	return fuel
}

func fuelForModules(m []int) int {
	var fuel int
	for _, f := range m {
		fuel = fuel + calcfuel(f)
	}
	return fuel
}

func fuel(m []int) int {
	var total int
	for _, f := range m {
		var fuel int
		fuel = f
		for {
			f := calcfuel(fuel)
			if f < 0 {
				break
			} else {
				fuel = f
				total = total + f
			}
		}
	}
	return total
}

func main() {
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalFuel int
	var m []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f, _ := strconv.Atoi(scanner.Text())
		m = append(m, f)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	totalFuel = fuelForModules(m)
	part2 := fuel(m)

	fmt.Printf("Total Fuel: %d\n", totalFuel)
	fmt.Printf("Part2 Total Fuel: %d\n", part2)
}
