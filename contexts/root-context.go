package contexts

import (
	"github.com/gocraft/web"
)

type RootContext struct{}

func (c *RootContext) Register(rw *web.ResponseWriter, request *web.Request) {

}

func (c *RootContext) Login(rw *web.ResponseWriter, request *web.Request) {

}

func (c *RootContext) Forgot(rw *web.ResponseWriter, request *web.Request) {

}

func (c *RootContext) Reset(rw *web.ResponseWriter, request *web.Request) {

}
