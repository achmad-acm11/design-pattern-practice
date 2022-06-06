package Chain_of_responsibility

import "strings"

type StringHandler interface {
	Process(s string) string
	SetNext(handler StringHandler)
}

type LowerCaseString struct {
	next StringHandler
}

func (l LowerCaseString) Process(s string) string {
	s = strings.ToLower(s)
	if l.next != nil {
		s = l.next.Process(s)
	}
	return s
}

func (l *LowerCaseString) SetNext(handler StringHandler) {
	l.next = handler
}

type RemoveSpace struct {
	next StringHandler
}

func (r RemoveSpace) Process(s string) string {
	s = strings.Replace(s, " ", "", -1)
	if r.next != nil {
		s = r.Process(s)
	}
	return s
}

func (r RemoveSpace) SetNext(handler StringHandler) {
	r.next = handler
}
