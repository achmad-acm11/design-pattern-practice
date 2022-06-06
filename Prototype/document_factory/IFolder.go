package document_factory

import "factory-pattern/Prototype"

type IFolder struct {
	Files []Prototype.INode
	name  string
}

func (I *IFolder) SetName(name string) {
	I.name = name
}
func (I IFolder) Print() string {
	return I.name
}

func (I IFolder) Clone() Prototype.INode {
	cloneFolder := &IFolder{name: I.name + "_clone"}
	var tmpFiles []Prototype.INode

	for _, file := range I.Files {
		tmpFiles = append(tmpFiles, file)
	}

	cloneFolder.Files = tmpFiles
	return cloneFolder
}
