package brbn

// TODO: come up with proper signature
type Interceptor func(*Context)

// TODO: add functions for adding to middleware
type Middleware []Interceptor
