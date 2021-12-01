package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_only_adds_first_value_when_arrow_points_to_beginning(t *testing.T) {
	var collection [2000]int

	pushValueIntoWindow(10, 0, &collection)

	assert.Equal(t, 10, collection[0])
	assert.Equal(t, 0, collection[1])
	assert.Equal(t, 0, collection[2])
}

func Test_only_adds_first_and_second_values_when_arrow_points_to_second_position(t *testing.T) {
	var collection [2000]int

	pushValueIntoWindow(10, 1, &collection)

	assert.Equal(t, 10, collection[0])
	assert.Equal(t, 10, collection[1])
	assert.Equal(t, 0, collection[2])
}

func Test_adds_all_three_values_when_arrow_points_is_greater_than_second_position(t *testing.T) {
	var collection [2000]int
	pushValueIntoWindow(199, 0, &collection)
	pushValueIntoWindow(200, 1, &collection)
	pushValueIntoWindow(208, 2, &collection)
	pushValueIntoWindow(210, 3, &collection)
	pushValueIntoWindow(200, 4, &collection)
	pushValueIntoWindow(207, 5, &collection)
	pushValueIntoWindow(240, 6, &collection)
	pushValueIntoWindow(269, 7, &collection)
	pushValueIntoWindow(260, 8, &collection)
	pushValueIntoWindow(263, 9, &collection)

	assert.Equal(t, 607, collection[0])
	assert.Equal(t, 618, collection[1])
	assert.Equal(t, 618, collection[2])
	assert.Equal(t, 617, collection[3])
	assert.Equal(t, 647, collection[4])
	assert.Equal(t, 716, collection[5])
	assert.Equal(t, 769, collection[6])
	assert.Equal(t, 792, collection[7])
}

func Test_should_count_number_of_increments(t *testing.T) {
	var collection [2000]int

	collection[0] = 607
	collection[1] = 618
	collection[2] = 618
	collection[3] = 617
	collection[4] = 647
	collection[5] = 716
	collection[6] = 769
	collection[7] = 792

	increments := countIncrements(collection)

	assert.Equal(t, 5, increments)
}
