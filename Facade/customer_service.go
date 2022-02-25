package Facade

type Ticket struct {
	Username string
	Type     string
	Message  string
}

type Response struct {
	Message string
}

func Handleticket(t Ticket) *Response {
	switch t.Type {
	case "food":
		response := HandleMessage(t.Message)
		return &Response{Message: response}
	}

	return &Response{
		Message: "Customer service is busy",
	}
}
