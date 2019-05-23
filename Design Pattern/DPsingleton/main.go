package main

import (
	"Dp/DPsingleton/singleton"
	"fmt"
)

func main() {
	s := singleton.New("lazyboy", 22, "China")
	s1 := singleton.New("asdf", 23, "Nibelungenlied")
	fmt.Println("s:  ", s.Name, s.Age, s.Country)
	fmt.Println("s1: ", s1.Name, s1.Age, s1.Country)

	s1.Name = "LazyboyChen7"
	fmt.Println("s:  ", s.Name, s.Age, s.Country)
	fmt.Println("s1: ", s1.Name, s1.Age, s1.Country)

	fmt.Println(&s, &s1)
	fmt.Println(&s.Name, &s.Age, &s.Country, &s.Extra, &s.Final)
	fmt.Println(&s1.Name, &s1.Age, &s1.Country, &s.Extra, &s.Final)
}
