package main

type submarine struct {
	horizontal int
	depth      int
}

func createSubmarine() submarine {
	return submarine{
		horizontal: 0,
		depth:      0,
	}
}

func (s *submarine) operate(operation operation) {
	switch operation.movement {
	case "forward":
		s.horizontal += operation.units
		break

	case "up":
		s.depth -= operation.units
		break

	case "down":
		s.depth += operation.units
		break
	}
}
