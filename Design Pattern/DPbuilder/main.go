package main

import (
	"Dp/DPbuilder/builder"
	"fmt"
)

func main() {
	manufacturingComplex := builder.ManufacturingDirector{}

	carBuilder := &builder.CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	fmt.Println(car)

	bikeBuilder := &builder.BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	fmt.Println(motorbike)
}
