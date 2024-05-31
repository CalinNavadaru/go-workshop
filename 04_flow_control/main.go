package main

import (
	"fmt"
	"strings"
)

func exemplu1() {
	arr := []int{100, 45, 67, 86, 12}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

	for idx, value := range arr {
		fmt.Println(idx, value)
	}
}

func exemplu2() {
	arr := []string{"Alex", "Andrei", "Alexandra", "Daniel", "Iulia", "Armand", "Alex", "Calin"}

	for _, value := range arr {
		fmt.Println(strings.Contains(value, "dr"))
	}
}

func exemplu3() {
	arr := []string{"Alex", "Andrei", "Alexandra", "Daniel", "Iulia", "Armand", "Alex", "Calin"}

	names := make(map[string]int)
	for _, value := range arr {
		names[value]++

		if _, ok := names[value]; ok {
			names[value]++
		} else {
			names[value] = 1
		}
	}
	fmt.Println(names)
}

func exemplu4() {
	var cond bool
	cond = true
	i := 0
	//for ;cond;
	for cond {
		i += 1
		if i == 5 {
			cond = false
		}
			
	}
}

func main() {
	exemplu1()
	exemplu2()
	exemplu3()
	exemplu4()
}
