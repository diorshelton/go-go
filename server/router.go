package server

type Router struct {
	Route []Route
}

type Route struct {
	Method  string
	Path    string
	Handler func(*Request) *Response
}
