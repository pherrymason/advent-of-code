package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_move_submarine_forward_straight(t *testing.T) {
	submarine := createSubmarine()
	forward := createForward(5)
	submarine.operate(forward)

	assert.Equal(t, 5, submarine.horizontal)
	assert.Equal(t, 0, submarine.depth)
	assert.Equal(t, 0, submarine.aim)
}

func Test_should_move_submarine_while_aiming(t *testing.T) {
	submarine := createSubmarine()
	submarine.operate(createDown(2))
	submarine.operate(createForward(5))

	assert.Equal(t, 5, submarine.horizontal)
	assert.Equal(t, 10, submarine.depth)
	assert.Equal(t, 2, submarine.aim)
}

func Test_should_aim_submarine_up(t *testing.T) {
	submarine := createSubmarine()

	submarine.operate(createDown(5))

	assert.Equal(t, 0, submarine.depth)
	assert.Equal(t, 0, submarine.horizontal)
	assert.Equal(t, 5, submarine.aim)
}

func Test_should_aim_submarine_down(t *testing.T) {
	submarine := createSubmarine()

	submarine.operate(createDown(5))
	submarine.operate(createUp(2))

	assert.Equal(t, 0, submarine.depth)
	assert.Equal(t, 0, submarine.horizontal)
	assert.Equal(t, 3, submarine.aim)
}
