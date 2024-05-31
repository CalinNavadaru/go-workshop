package main

import "fmt"

//use pointers if your struct is huge
//use pointers if you want to mutate your variable
//use pointers if you want to share your variabless accross functions

//don't use pointers for things that are pointers under the hood, like slices

func changeByVal(n int) {
	n += 1
}

func changeByPPtr(n *int) {
	*n += 1
}

func exercise1() {
	n := 10

	changeByVal(n)
	fmt.Println(n)

	changeByPPtr(&n)
	fmt.Println(n)
}

type SomeStruct struct {
	Name string;
	Age int;

}

func (s SomeStruct) changeVal() {
	s.Age = 10
}

func (s *SomeStruct) changePtr() {
	s.Age = 10
}

func exercise2() {
	s := SomeStruct{
		Name: "Marian",
		Age: 23,
	}

	s.changeVal()
	fmt.Println(s)
	s.changePtr()
	fmt.Println(s)
}

func main() {
	exercise1()
	exercise2()
}