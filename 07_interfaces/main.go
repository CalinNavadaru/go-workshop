package main

import "fmt"

type HelloWorlder interface {
	HelloWorld()
}

type RomanianHelloWorder struct {

}

type EnglishHelloWorder struct {

}

func (r RomanianHelloWorder) HelloWorld() {
	fmt.Println("Messi")
}


func (r EnglishHelloWorder) HelloWorld() {
	fmt.Println("Ronaldo")
}

func (r RomanianHelloWorder) Cefaci() {
	fmt.Println("Ce faci domne?")
}

func main() {
	var says HelloWorlder

	says = RomanianHelloWorder{}
	says.HelloWorld()

	//type assertion
	romanian, ok := says.(RomanianHelloWorder)
	fmt.Println(romanian, ok)
	romanian.Cefaci()

	says = EnglishHelloWorder{}
	says.HelloWorld()

	switch says.(type) {
	case EnglishHelloWorder:
		fmt.Println("E englez ba")
	case RomanianHelloWorder:
		fmt.Println("E de-al nostru ")
	}
}