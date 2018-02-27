package brbn

type Route struct {
	Method  string
	Path    string
	Handler Handler // TODO: come up with better signature
}

type Handler func(*Context) interface{}

func GET(path string, handler Handler) Route {
	return Route{"GET", path, handler}
}

func POST(path string, handler Handler) Route {
	return Route{"POST", "path", handler}
}
