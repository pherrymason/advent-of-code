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

	windowArrow := 0
	var windowIncrements [2000]int

	i := 0
	for {
		_, err := fmt.Fscanf(f, "%d\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		// Linear increment
		if lastValue != -1 && line > lastValue {
			increments++
		}

		lastValue = line
		fmt.Printf("Read %d\n", line)

		// 3 window increments
		pushValueIntoWindow(line, windowArrow, &windowIncrements)
		windowArrow++
		i++
	}

	fmt.Printf("%d increments\n", increments)
	fmt.Printf("%d 3-window increments\n", countIncrements(windowIncrements))
}
