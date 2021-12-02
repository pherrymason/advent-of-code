package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_move_submarine_forward(t *testing.T) {
	submarine := createSubmarine()
	forward := createOperation("forward", 5)
	submarine.operate(forward)

	assert.Equal(t, 5, submarine.horizontal)
	assert.Equal(t, 0, submarine.depth)
}

func Test_should_move_submarine_deeper(t *testing.T) {
	submarine := createSubmarine()
	down := createOperation("down", 5)
	submarine.operate(down)

	assert.Equal(t, 5, submarine.depth)
	assert.Equal(t, 0, submarine.horizontal)
}

func Test_should_move_submarine_up(t *testing.T) {
	submarine := createSubmarine()
	down := createOperation("down", 5)
	submarine.operate(down)

	submarine.operate(createOperation("up", 2))

	assert.Equal(t, 3, submarine.depth)
	assert.Equal(t, 0, submarine.horizontal)
}
