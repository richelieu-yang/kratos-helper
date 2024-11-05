package kpprofKit

import (
	_ "net/http/pprof"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"net/http"
	"net/http/pprof"
)

func init() {
	/*
		TODO: 可能会导致bug: 在 http.DefaultServeMux 上自行绑定的路由丢了.

		net/http/pprof的 init() 会自动注册一些路由到默认的 HTTP 路由器中.
		此处代码的目的: 取消这些自动绑定的路由.
	*/
	http.DefaultServeMux = http.NewServeMux()
}

// RegisterPprof 注册pprof相关的路由.
/*
参考了官方example: https://github.com/go-kratos/examples/blob/main/http/pprof/main.go
*/
func RegisterPprof(s *khttp.Server) {
	s.HandleFunc("/debug/pprof", pprof.Index)
	s.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.HandleFunc("/debug/pprof/trace", pprof.Trace)
	s.HandleFunc("/debug/allocs", pprof.Handler("allocs").ServeHTTP)
	s.HandleFunc("/debug/block", pprof.Handler("block").ServeHTTP)
	s.HandleFunc("/debug/goroutine", pprof.Handler("goroutine").ServeHTTP)
	s.HandleFunc("/debug/heap", pprof.Handler("heap").ServeHTTP)
	s.HandleFunc("/debug/mutex", pprof.Handler("mutex").ServeHTTP)
	s.HandleFunc("/debug/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
}
