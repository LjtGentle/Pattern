package main

import "fmt"

func main() {
	normalBuilder := getBuilder(NormalType)
	iglooBuilder := getBuilder(IglooType)

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	printDetails(normalHouse)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	printDetails(iglooHouse)
}

func printDetails(h house) {
	fmt.Println("house door type: ",h.doorType)
	fmt.Println("house window type: ",h.windowType)
	fmt.Println("house num floor: ",h.floor)
}