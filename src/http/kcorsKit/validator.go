package kcorsKit

import "strings"

type (
	Validator interface {
		ValidateOrigin(origin string) bool
	}

	validatorImpl struct {
		allowAll bool

		allowOrigins    []string
		wildcardOrigins [][]string
	}
)

func (impl *validatorImpl) ValidateOrigin(origin string) bool {
	if origin == "" {
		return true
	}

	for _, value := range impl.allowOrigins {
		if value == origin {
			return true
		}
	}

	/* 域名匹配 */
	for _, w := range impl.wildcardOrigins {
		if w[0] == "*" && strings.HasSuffix(origin, w[1]) {
			return true
		}
		if w[1] == "*" && strings.HasPrefix(origin, w[0]) {
			return true
		}
		if strings.HasPrefix(origin, w[0]) && strings.HasSuffix(origin, w[1]) {
			return true
		}
	}

	return false
}
