package piscine

//RecursivePower will print the power of a number nb in a recursive way
func RecursivePower(nb int, power int) int {

	if power < 0 || power > 100 && nb == 0 {

		return 0
	} else if power == 0 && nb >= 0 || nb < 0 && power == 1 {

		return 1
	}

	return nb * RecursivePower(nb, power-1)
}
