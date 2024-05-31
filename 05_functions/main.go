package main


func division(divident int, divisor int) (int, int) {
	q := divident / divisor
	r := divident % divisor
	return q, r
}

func division2(divident int, divisor int) (q int, r int) {
	q = divident / divisor
	r = divident % divisor
	return q, r
}

type dividefunc func(int, int) (int, int)

type Coordinate [2]int
func (c Coordinate) SumOfCoords() int {
	return c[1] + c[0]
}

func main() {
	division(5, 3)
	division2(5, 3)
}