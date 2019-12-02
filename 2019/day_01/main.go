package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int64

	for scanner.Scan() {
		mod, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Printf("Error converting string to int: %v", err)
		}
		fuel := calcFuel(int64(mod), 0)
		sum += fuel
	}
	log.Printf("Got Sum: %v", sum)
}

func calcFuel(fuel, sum int64) int64 {
	for fuel >= 0 {
		fuel = (int64(fuel) / 3) - 2
		if fuel <= 0 {
			break
		}
		sum += fuel
		calcFuel(fuel, sum)
	}
	return sum
}
