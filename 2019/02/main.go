package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	seq := getFileContents() 
		for k := 1; k < 100; k++ {
			for j := 1; j < 100; j++ {
				newSeq := runSequence(seq, k, j)
				if newSeq[0] == 19690720 {
					log.Printf("Noun: %v Verb: %v", k, j)
					break
				}
			}
	}
}

func getFileContents() []int32 {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sequence []int32
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), ",")

		for _, s := range str {
			integer, err := strconv.Atoi(s)
			if err != nil {
				log.Printf("Error converting integer: %v", err)
			}
			sequence = append(sequence, int32(integer))
		}
	}
	return sequence
}


func copyArray(array []int32) []int32 {
	arrNew := make([]int32, 0)
	arrNew = append(arrNew, array...)
	return arrNew
}

func runSequence(pro []int32, noun, verb int) []int32 {
	seq := copyArray(pro)
	seq[1] = int32(noun)
	seq[2] = int32(verb)

	for i := 0; seq[i] != 99; i++ {
		res := seq[i+3]
		pos1 := seq[i+1]
		pos2 := seq[i+2]
		if seq[i] == 1 && seq[i+3] <= int32(len(seq)) && seq[i+1] <= int32(len(seq)) && seq[i+2] <= int32(len(seq)){
			seq[res] = (seq[pos1] + seq[pos2])
			i += 3
		}
		if seq[i] == 2 && seq[i+3] <= int32(len(seq)) && seq[i+1] <= int32(len(seq)) && seq[i+2] <= int32(len(seq)){
			seq[res] = (seq[pos1] * seq[pos2])
			i += 3
		}
	}
	return seq
}