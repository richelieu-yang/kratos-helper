package kcorsKit

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	ghttp "net/http"
)

func NewCorsServerOption(allowOrigins []string) http.ServerOption {
	filterFunc := NewCorsFilterFunc(allowOrigins)
	return http.Filter(filterFunc)
}

// NewCorsFilterFunc
/*
TODO: github.com/gorilla/handlers v1.5.2 有点问题: 当收到预检请求且发现跨域，返回的状态码是 200 而非 403，虽然跨域了但真正的handler不会触发，但无伤大雅，后续看官方会不会处理 || 自己fork.

@param allowOrigins: 允许的跨域域名 	(1) 支持域名匹配
									(2) 可以为nil，即都允许（并非通过 "*" 来实现）
									(3) 如果有一元素为"*"，即都允许（并非通过 "*" 来实现）
*/
func NewCorsFilterFunc(allowOrigins []string) http.FilterFunc {
	validator := NewValidator(allowOrigins)

	return handlers.CORS(
		handlers.AllowedOriginValidator(func(s string) bool {
			return validator.ValidateOrigin(s)
		}),
		// 设置允许的 HTTP 方法
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		// 设置允许的请求头
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),

		/*
			缓存时间 (Access-Control-Max-Age): 设置浏览器可以缓存此预检请求的时长，以减少频繁的预检请求。
			e.g.Access-Control-Max-Age: 86400（即 24 小时）
		*/
		handlers.MaxAge(86400),
		handlers.OptionStatusCode(ghttp.StatusNoContent),
	)
}
