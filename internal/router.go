package server

import (
	"fmt"
	"strings"
)

type route struct {
	method string
	path string
	handler func(ctx Context) string
}

type Handlers struct {
	handler func(ctx Context) string
	params map[string]string
}

type RouterType struct {
	Routes []route
}

type Router interface {
	AddRoute(method string, path string, handler func(ctx Context) string)
	MatchRoute(method string, path string) []Handlers
	GET(path string, handler func(ctx Context) string)
	POST(path string, handler func(ctx Context) string)
}

func NewRouter() Router {
	return &RouterType{}
}

func (r *RouterType) AddRoute(method string, path string, handler func(ctx Context) string) {
	route := route{
		method: method,
		path: path,
		handler: handler,
	}
	r.Routes = append(r.Routes, route)
}

func (r *RouterType) MatchRoute(method string, path string) []Handlers {
	var matchedRoutes []Handlers
	for _, route := range r.Routes {
		if route.method != method {
			continue
		}
		if route.path == "/*" || route.path == "*" {
			matchedRoutes = append(matchedRoutes, Handlers{handler: route.handler})
			continue
		}
		hasStar := strings.Contains(route.path, "*")
		hasNamedParam := strings.Contains(route.path, ":")
		if !hasStar && !hasNamedParam {
			if route.path == path || route.path + "/" == path {
				matchedRoutes = append(matchedRoutes, Handlers{handler: route.handler})
			}
		} else if hasStar {
			if strings.HasPrefix(path, strings.TrimRight(route.path, "*")) {
				matchedRoutes = append(matchedRoutes, Handlers{handler: route.handler})
			}
		} else if hasNamedParam {
			pathParts := strings.Split(path, "/")
			routeParts := strings.Split(route.path, "/")
			if len(pathParts) != len(routeParts) {
				fmt.Println("Path parts not equal")
				continue
			}
			params := make(map[string]string)
			for i, part := range routeParts {
				if strings.HasPrefix(part, ":") {
					params[part[1:]] = pathParts[i]
				} else if part != pathParts[i] {
					break
				}
				if i == len(routeParts) - 1 {
					matchedRoutes = append(matchedRoutes, Handlers{handler: route.handler, params: params})
				}
			}
		}
	}
	return matchedRoutes
}

func (r *RouterType) GET(path string, handler func(ctx Context) string) {
	r.AddRoute("GET", path, handler)
}

func (r *RouterType) POST(path string, handler func(ctx Context) string) {
	r.AddRoute("POST", path, handler)
}
