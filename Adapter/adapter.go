package Adapter

import "fmt"

type AdapterAbstract interface {
	Play(m Mp3)
}

type Mp3 struct {
	Lagu []byte
}

type Kaset struct {
	Lagu string
}

type Walkman struct {
}

func (w Walkman) PlayMusic(k Kaset) {
	fmt.Println(k.Lagu)
}

type Mp3toKasetAdapter struct {
	Adaptee Walkman
}

func (k Mp3toKasetAdapter) Play(m Mp3) {
	objKaset := Kaset{
		Lagu: string(m.Lagu),
	}

	k.Adaptee.PlayMusic(objKaset)
}
