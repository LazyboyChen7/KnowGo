package main

import (
	"Dp/DPfactory/factory"
	"fmt"
)

func main() {
	f := new(factory.Factory)
	p := f.Create("a")
	fmt.Println(p.GetName())
	p = f.Create("b")
	fmt.Println(p.GetName())
}
