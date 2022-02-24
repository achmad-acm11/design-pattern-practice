package main

import "factory-pattern/Adapter"

func main() {
	//var product Factory.Product
	//product = &Factory.Car{}
	//
	//typeObj := product.DoStuff()
	//typeObj.Running()
	//
	//product = &Factory.Boat{}
	//typeObj = product.DoStuff()
	//typeObj.Running()

	mp3 := Adapter.Mp3{
		Lagu: []byte("Lagu Baru"),
	}

	var ad Adapter.AdapterAbstract
	ad = Adapter.Mp3toKasetAdapter{}

	ad.Play(mp3)
}
