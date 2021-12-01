package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	var line int
	lastValue := -1
	increments := 0
	for {
		_, err := fmt.Fscanf(f, "%d\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		if lastValue != -1 && line > lastValue {
			increments++
		}

		lastValue = line
		fmt.Printf("Read %d\n", line)
	}

	fmt.Printf("%d increments\n", increments)
}
