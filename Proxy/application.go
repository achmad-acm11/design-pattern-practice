package Proxy

type Application struct {
}

func (a *Application) HandleRequest(url string, method string) (int, string) {
	if url == "api/status" && method == "GET" {
		return 200, "OK"
	}
	if url == "api/create/user" && method == "POST" {
		return 201, "Created"
	}
	return 404, "Not found"
}
