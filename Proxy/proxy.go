package Proxy

type nginx struct {
	application    Application
	maxAllowAccess int
	rateLimiter    map[string]int
}

func NewNginxServer() *nginx {
	return &nginx{
		application:    Application{},
		maxAllowAccess: 2,
		rateLimiter:    make(map[string]int),
	}
}

func (n *nginx) HandleRequest(url string, method string) (int, string) {
	allowed := n.checkRateLimit(url)

	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

func (n *nginx) checkRateLimit(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowAccess {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
