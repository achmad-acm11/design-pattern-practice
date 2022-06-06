package Composite

import "fmt"

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(s string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s \n", s, f.Name)
	for _, c := range f.Components {
		c.Search(s)
	}
}

func (f *Folder) AddComponent(c Component) {
	f.Components = append(f.Components, c)
}
