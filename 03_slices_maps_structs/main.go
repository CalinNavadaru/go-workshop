package main

import "fmt"

//arrays, slices, maps
//structs

func example1() {
	var arr [5] int

	fmt.Println(arr)

	arr[1] = 10

	fmt.Println(arr)

	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}

func example2() {
	arr := make([]int, 0, 10)
	arr2 := make([]int, 5, 10)
	arr3 := make([]int, 5) /// <=> arr3 := make([]int, 5, 5)
	arr4 := []int{1, 2, 3, 4}

	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	fmt.Println(len(arr2))
	fmt.Println(cap(arr2))

	fmt.Println(len(arr3))
	fmt.Println(cap(arr3))

	fmt.Println(len(arr4))
	fmt.Println(cap(arr4))
}

func example3() {
	arr := []int{1, 2, 3, 4}

	arr = append(arr, 4, 5, 6)

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	var arr2 []int
	fmt.Println(arr2 == nil)

	var arr3 = make([]int, 0)
	fmt.Println(arr3 == nil)
}

func example4() {
	mymap := make(map[int]int, 0)

	mymap[10] = 9
	mymap[9] = 10

	fmt.Println(mymap)

	delete(mymap, 10)

	fmt.Println(mymap)
}

func example5() {
	set := make(map[int]struct{})
	set[10] = struct{}{}
	delete(set, 10) 

}

type SomeStruct struct {
	Name string;
	Age int;
}

func example6() {
	s := SomeStruct{
		Name: "Marius",
		Age: 32,
	}
	s1 := SomeStruct{}
	fmt.Printf("%v, %v \n", s, s1)
}

type Struct2 struct {
	Job string;
	SomeStruct;
}

func example7() {
	s := SomeStruct{
		Name: "Marius",
		Age: 32,
	}
	s1 := Struct2{
		Job: "Soldat",
		SomeStruct: s,
	}
	fmt.Printf("%v, %v \n", s, s1)
	fmt.Println(s1.SomeStruct.Age)
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}