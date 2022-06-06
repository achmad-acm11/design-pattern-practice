package Proxy

type Server interface {
	HandleRequest(string, string) (int, string)
}
