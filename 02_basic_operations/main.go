package main

import (
	"fmt"
	"strconv"
)

//int - 32/64
//float - 64
//string
//bool

///1454 + 234 - 345

///go doc
func execitiu1a() {
	fmt.Println("Merge problema 1a) :)")
	fmt.Println(1454 + 234 - 345)
	var x1, x2, x3 int
	x1 = 1454
	x2 = 234
	x3 = 345

	fmt.Println(x1 + x2 - x3)

	x4 := 1454
	x5 := 234
	x6 := 345

	fmt.Println(x4 + x5 - x6)

	var z1, z2 float64
	z1 = 3
	z2 = 3.13

	fmt.Println(z1 / z2)

	a1 := 10e0
	a2 := 2.5

	fmt.Println(a1 / a2)
}

func execitiu2a() {
	var x1, x2, x3 bool
	x1 = true
	x2 = false
	x3 = true

	fmt.Println(x1 && x2 || x3)
}

func exercitiu3a() {
	alex := "alex"
	armand := "armand"
	fmt.Println(alex + armand)

	a := 'a';
	b := 'b';
	fmt.Println(a + b)

	fmt.Println(alex[0:3])
}

func exercitiu4a() {
	x := 10
	xstr := strconv.Itoa(x)

	fmt.Println(xstr)

	y, _ := strconv.Atoi(xstr)

	fmt.Println(y)
}

func main() {
	execitiu1a()
	execitiu2a()
	exercitiu3a()
	exercitiu4a()
}
