package main

import (
	"fmt"
	"slices"
)

/*
-Adding a new programming
-Listing all programmers (as an array)
-Send the top programmer to work ( based on language and yoe constraints)
-- programmer needs to work in that language
-- programmer needs to have at least that many yoe
*/

type Programmer struct {
	Name string;
	Language string;
	Yoe int;
}

type BenchPool interface {
	AddProgrammer(p Programmer)
	ListAllProgrammers()
	SendToWork(language string, minYears int) Programmer
}

type ArrayBenchPool struct {
	programmers []Programmer
}

func (a *ArrayBenchPool) AddProgrammer(p Programmer) {
	a.programmers = append(a.programmers, p)
}

func (a *ArrayBenchPool) ListAllProgrammers() []Programmer {
	fmt.Println(a.programmers)
	return slices.Clone(a.programmers)
}

func main() {

}