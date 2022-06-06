package document_factory

import "factory-pattern/Prototype"

type IFile struct {
	Name string
}

func (I IFile) Print() string {
	return I.Name
}

func (I IFile) Clone() Prototype.INode {
	return &IFile{
		Name: I.Name + "_clone",
	}
}
