package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	geo "github.com/paulmach/go.geo"
)

type Coords struct {
	Point []geo.Point
}

func main() {
	
	test1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	test2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	// commands := getFileContents()

	// one := getCoords(commands[0][0])
	// two := getCoords(commands[1][0])

	one := getCoords(test1)
	two := getCoords(test2)

	log.Printf("one: %v \n two: %v", one, two)

	find := Intersection(one.Point, two.Point)

	log.Printf("find: %v", find)
	// log.Printf("One: %v, \n Two: %v", one, two)
	
}

func getCoords(commands string) Coords {
	coords := Coords{}
	c := strings.Split(commands, ",")
	pt := []float64{0, 0}
	for _, s := range c {
		split := strings.Split(s, "")
		cmd := split[0]
		steps, err := strconv.ParseFloat((strings.Join(split[1:], "")), 64)
		if err != nil {
			log.Printf("Error getting int from string: %v", err)
		}

		switch cmd {
		case "R":
			pt[0] += steps
		case "L":
			pt[0] -= steps
		case "U":
			pt[1] += steps
		case "D":
			pt[1] -= steps
		}
		coords.Point = append(coords.Point, *geo.NewPoint(pt[0], pt[1]))
	}
	return coords
}

func getFileContents() [][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var commands [][]string

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), "\n")
		commands = append(commands, str)
	}
	return commands
}


func Intersection(a, b []geo.Point) (c []geo.Point) {
	m := make(map[geo.Point]bool)

	for _, item := range a {
			m[item] = true
	}

	for _, item := range b {
			if _, ok := m[item]; ok {
					c = append(c, item)
			}
	}
	return
}