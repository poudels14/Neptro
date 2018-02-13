package router

type RouteStruct struct {
	Method     string
	Path       string
	Controller interface{}
}

func Route(method string, path string, controller interface{}) RouteStruct {
	return RouteStruct{method, path, controller}
}
