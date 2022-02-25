package main

import (
	"factory-pattern/Proxy"
	"fmt"
	"time"
)

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

	//mp3 := Adapter.Mp3{
	//	Lagu: []byte("Lagu Baru"),
	//}
	//
	//var ad Adapter.AdapterAbstract
	//ad = Adapter.Mp3toKasetAdapter{}
	//
	//ad.Play(mp3)

	youtubeImp := Proxy.YoutubeImpl{}

	youtubeCache := Proxy.CacheYoutube{
		YoutubeService: youtubeImp,
	}

	go youtubeCache.Run()

	fmt.Println(youtubeCache.ListVideo())
	fmt.Println(youtubeCache.ListVideo())

	<-time.After(7 * time.Second)

	fmt.Println(youtubeCache.ListVideo())
}
