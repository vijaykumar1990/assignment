package http

import (
	"github.com/labstack/echo"
)

type Router interface {
	GET(url string, f func(ctx echo.Context) error)
	POST(url string, f func(ctx echo.Context) error)
	DELETE(url string, f func(ctx echo.Context) error)
	SERVE(port string)
}

type echoRouter struct {
}

func NewEchoRouter() Router {
	return &echoRouter{}
}

var echoDispatcher = echo.New()

func (*echoRouter) GET(url string, f func(ctx echo.Context) error) {
	echoDispatcher.GET(url, f)
}

func (*echoRouter) POST(url string, f func(ctx echo.Context) error) {
	echoDispatcher.POST(url, f)
}
func (*echoRouter) DELETE(url string, f func(ctx echo.Context) error) {
	echoDispatcher.DELETE(url, f)
}
func (*echoRouter) SERVE(port string) {
	echoDispatcher.Start(port)
}
