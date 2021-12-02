package main

import (
	"fmt"
	"io"
	"os"
)

func openInput() *os.File {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	return f
}

func readNextMovement(f *os.File) operation {
	var movement string
	var units int
	_, err := fmt.Fscanf(f, "%s %d\n", &movement, &units)
	if err != nil {
		if err == io.EOF {
			return createEndOperation()
		}
	}

	return createOperation(movement, units)
}

func main() {
	submarine := createSubmarine()

	f := openInput()
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	for {
		operation := readNextMovement(f)
		if operation.isEnd() {
			break
		}
		submarine.operate(operation)
		fmt.Printf("%s %d\n", operation.movement, operation.units)
	}

	fmt.Printf("Submarine horizontal: %d, depth: %d, aim: %d => %d",
		submarine.horizontal,
		submarine.depth,
		submarine.aim,
		submarine.horizontal*submarine.depth)
}
