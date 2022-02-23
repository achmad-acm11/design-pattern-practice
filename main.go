package main

import (
	"factory-pattern/Factory"
)

func main() {
	var product Factory.Product
	product = &Factory.Car{}

	typeObj := product.DoStuff()
	typeObj.Running()

	product = &Factory.Boat{}
	typeObj = product.DoStuff()
	typeObj.Running()
}
