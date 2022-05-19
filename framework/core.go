package framework

import (
	"fmt"
	"net/http"
)

type Core struct{}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Println("here is core")
}
