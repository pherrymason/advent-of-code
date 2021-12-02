package main

type submarine struct {
	horizontal int
	depth      int
	aim        int
}

func createSubmarine() submarine {
	return submarine{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}
}

func (s *submarine) operate(operation operation) {
	switch operation.movement {
	case "forward":
		s.horizontal += operation.units
		s.depth += s.aim * operation.units
		break

	case "up":
		s.aim -= operation.units
		break

	case "down":
		s.aim += operation.units
		break
	}
}
