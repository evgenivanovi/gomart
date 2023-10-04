package infra

import "net/http"

/* __________________________________________________ */

type Controller interface {
	Handle(writer http.ResponseWriter, request *http.Request)
}

func AsHandlerFunc(c Controller) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		c.Handle(writer, request)
	}
}

/* __________________________________________________ */
