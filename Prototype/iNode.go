package Prototype

type INode interface {
	Print() string
	Clone() INode
}
