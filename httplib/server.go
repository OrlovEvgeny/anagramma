package httplib

import (
	"anagramma"
	"anagramma/bootstrap"
	"anagramma/httplib/controller"
	"context"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

//HttpServer
func HttpServer(ctx context.Context, addr string) {
	ctrl := new(controller.Controller)
	ctrl.HashMapAlgorithm = ctx.Value(bootstrap.HashMapAlgorithm).(*anagramma.HashMap)

	s := &fasthttp.Server{
		Name:               "Anagramma service",
		Handler:            routeLoader(ctrl).Handler,
		MaxRequestBodySize: 50 * 1024 * 1024,
	}

	log.Printf("http_server start by %s\n", addr)
	err := s.ListenAndServe(addr)
	if err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

//routeLoader
func routeLoader(ctrl *controller.Controller) *fasthttprouter.Router {
	router := fasthttprouter.New()

	/**
	* Main ctrl
	* curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'
	 */
	router.POST("/load", ctrl.Load)

	/**
	* Main ctrl
	* curl 'localhost:8080/get?word=foobar'  => ["foobar","boofar"]
	 */
	router.GET("/get", ctrl.Get)

	return router
}
