package smpkg

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RouterService struct {
	router      http.Handler
	middlewares []func(http.Handler) http.Handler
}

func NewRouter() *RouterService {
	return &RouterService{
		router: httprouter.New(),
	}
}

func (r *RouterService) Get(path string, handle httprouter.Handle) {
	r.router.(*httprouter.Router).GET(path, handle)
}

func (r *RouterService) Post(path string, handle httprouter.Handle) {
	r.router.(*httprouter.Router).POST(path, handle)
}

func (r *RouterService) UseMiddleware(md func(http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, md)
}

func (r *RouterService) UseCorsAccess() *RouterService {
	r.router.(*httprouter.Router).GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	return r
}
