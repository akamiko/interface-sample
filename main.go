package main

import (
	"fmt"
)

// 人間
type Human interface {
	Greeting()
}

// 日本人
type Japanese struct{}

func (j *Japanese) Greeting() {
	fmt.Println("こんにちは!")
}

// アメリカ人
type American struct{}

func (a *American) Greeting() {
	fmt.Println("hello!")
}

func HumanGreeting(h Human) {
	h.Greeting()
}

func main() {
	j := new(Japanese)
	HumanGreeting(j)

	a := new(American)
	HumanGreeting(a)
}
