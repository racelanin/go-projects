package piscine

func RockAndRoll(n int) string {
	switch {
	case n < 0:
		return "error: number is negative\n"
	case n%2 == 0 && n%3 == 0:
		return "rock and roll\n"
	case n%3 == 0:
		return "roll\n"
	case n%2 == 0:
		return "rock\n"

	}
	return "error: non divisible\n"
}
