package kcorsKit

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
