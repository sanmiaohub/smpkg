package smpkg

import "net/http"

// EngineProvide 封装的服务提供者, 在实际的router路由控制器处理业务之前，可以传递中间件处理者
type EngineProvide struct {
	router RouterService
}

func NewEngine(router RouterService) *EngineProvide {
	return &EngineProvide{
		router: router,
	}
}

// ServeHTTP 先处理中间件, 在处理路由业务
func (ep *EngineProvide) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 对ServeHTTP的调用要放后面, 不然先调用的话，是从里到外进行的, 中间件是需要从外到里走的
	next := ep.router.router
	// middleware1(middleware2(middleware3(router)))
	if len(ep.router.middlewares) > 0 {
		for i := range ep.router.middlewares {
			// 反向包裹的, 在前面的中间件, 越往后执行
			process := ep.router.middlewares[len(ep.router.middlewares)-1-i]
			next = process(next)
		}
	}
	next.ServeHTTP(w, r)
}
