package server

import (
	"strings"
)

type Router struct {
	Routes []Route
}

type Route struct {
	Method  string
	Path    string
	Handler func(*Request) *Response
}

func (router *Router) Register(method, path string, handler func(*Request) *Response) {
	route := Route{Method: method, Path: path, Handler: handler}
	router.Routes = append(router.Routes, route)
}

func (router *Router) Match(request *Request) (func(*Request) *Response, map[string]string) {

outer:
	for _, route := range router.Routes {

		if route.Method != request.Method {
			continue
		}

		requestStr := strings.TrimLeft(request.URL, "/")
		routeStr := strings.TrimLeft(route.Path, "/")

		requestSegments := strings.Split(requestStr, "/")
		routeSegments := strings.Split(routeStr, "/")

		if len(requestSegments) != len(routeSegments) {
			continue
		}

		paramsMap := make(map[string]string)

		for i, routeSeg := range routeSegments {
			requestSeg := requestSegments[i]

			if strings.HasPrefix(routeSeg, ":") {
				//the route is dynamic
				paramsMap[routeSeg[1:]] = requestSeg
			} else {
				// the route is static and segments must be equal
				if routeSeg != requestSeg {
					continue outer
				}
			}
		}
		return route.Handler, paramsMap
	}

	return nil, nil
}
