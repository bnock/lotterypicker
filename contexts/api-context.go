package contexts

import (
	"github.com/bradmang18/lotterypicker/models"
	"github.com/gocraft/web"
)

type APIContext struct {
	User models.User
}

/**
 * Actions
 */
func (c *APIContext) List(rw *web.ResponseWriter, request *web.Request) {

}

func (c *APIContext) Create(rw *web.ResponseWriter, request *web.Request) {

}

func (c *APIContext) Find(rw *web.ResponseWriter, request *web.Request) {

}

func (c *APIContext) Update(rw *web.ResponseWriter, request *web.Request) {

}

func (c *APIContext) Delete(rw *web.ResponseWriter, request *web.Request) {

}

/**
 * Middleware
 */
func (c *APIContext) UserRequired(rw web.ResponseWriter, request *web.Request, next web.NextMiddlewareFunc) {
	next(rw, request)
}
