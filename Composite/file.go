package Composite

import "fmt"

type File struct {
	Name string
}

func (f *File) getName() string {
	return f.Name
}

func (f *File) Search(s string) {
	fmt.Printf("Searching for keyword %s in file %s \n", s, f.Name)
}
