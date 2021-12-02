package main

type operation struct {
	movement string
	units    int
	end      bool
}

func createEndOperation() operation {
	return operation{
		movement: "final",
		units:    0,
		end:      true,
	}
}

func createOperation(movement string, units int) operation {
	return operation{
		movement: movement,
		units:    units,
		end:      false,
	}
}

func (o *operation) isEnd() bool {
	return o.end
}
