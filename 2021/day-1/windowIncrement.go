package main

func pushValueIntoWindow(value int, windowArrow int, windowIncrements *[2000]int) {

	windowIncrements[windowArrow] += value

	if windowArrow > 0 {
		// Increment windowArrow -1
		windowIncrements[windowArrow-1] += value
	}

	if windowArrow > 1 {
		// Increment windowArrow -2
		windowIncrements[windowArrow-2] += value
	}
}

func countIncrements(windowIncrements [2000]int) int {
	increments := 0
	for j := 1; j < 2000; j++ {
		if windowIncrements[j] > windowIncrements[j-1] {
			increments++
		}
	}
	return increments
}
