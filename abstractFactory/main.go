package main

import "fmt"

func main() {
	adidasFactory,_ := getSportsFactory(Adidas)
	nikeFactory,_ := getSportsFactory(Nike)

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShoeDetails(adidasShoe)

	printShirtDetails(nikeShirt)
	printShirtDetails(adidasShirt)
}


func printShoeDetails(s iShoe) {
	fmt.Println("Logo: ",s.getLogo())
	fmt.Println("Size: ",s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Println("Logo: ",s.getLogo())
	fmt.Println("Size: ",s.getSize())
	fmt.Println()
}