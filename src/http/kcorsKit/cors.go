package kcorsKit

import "github.com/go-kratos/kratos/v2/transport/http"
import "github.com/gorilla/handlers"

func NewCorsServerOption(allowOrigins []string) http.ServerOption {
	filterFunc := NewCorsFilterFunc(allowOrigins)
	return http.Filter(filterFunc)
}

// NewCorsFilterFunc
/*
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
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)
}
