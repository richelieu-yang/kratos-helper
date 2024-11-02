package kcorsKit

import (
	"errors"
	"strings"
)

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

func NewValidator(allowOrigins []string) Validator {
	allowOrigins = normalize(allowOrigins)

	allowAll := false
	if len(allowOrigins) == 0 {
		allowAll = true
	} else {
		for _, origin := range allowOrigins {
			if origin == "*" {
				allowAll = true
				break
			}
		}
	}
	if allowAll {
		return &validatorImpl{
			allowAll: true,
		}
	}

	return &validatorImpl{
		allowAll:        false,
		allowOrigins:    allowOrigins,
		wildcardOrigins: parseWildcardRules(allowOrigins),
	}
}

func (impl *validatorImpl) ValidateOrigin(origin string) bool {
	if impl.allowAll || origin == "" {
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

/*
copy from github.com/gin-contrib/cors v1.7.2

效果: 去重、英文字母小写.
*/
func normalize(values []string) []string {
	if values == nil {
		return nil
	}
	distinctMap := make(map[string]bool, len(values))
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)

		/* Richelieu: 新增"去空" */
		if value == "" {
			continue
		}

		value = strings.ToLower(value)
		if _, seen := distinctMap[value]; !seen {
			normalized = append(normalized, value)
			distinctMap[value] = true
		}
	}
	return normalized
}

/*
copy from github.com/gin-contrib/cors v1.7.2
*/
func parseWildcardRules(allowOrigins []string) [][]string {
	var wRules [][]string

	for _, o := range allowOrigins {
		if !strings.Contains(o, "*") {
			continue
		}

		if c := strings.Count(o, "*"); c > 1 {
			panic(errors.New("only one * is allowed").Error())
		}

		i := strings.Index(o, "*")
		if i == 0 {
			wRules = append(wRules, []string{"*", o[1:]})
			continue
		}
		if i == (len(o) - 1) {
			wRules = append(wRules, []string{o[:i], "*"})
			continue
		}

		wRules = append(wRules, []string{o[:i], o[i+1:]})
	}

	return wRules
}
