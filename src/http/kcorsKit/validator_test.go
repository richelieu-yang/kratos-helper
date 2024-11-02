package kcorsKit

import (
	"fmt"
	"testing"
)

func TestNewValidator(t *testing.T) {
	{
		v := NewValidator([]string{})
		fmt.Println(v.ValidateOrigin("https://www.baidu.com")) // true
	}

	{
		v := NewValidator([]string{"", ""})
		fmt.Println(v.ValidateOrigin("https://www.baidu.com")) // true
	}

	{
		v := NewValidator([]string{"*.baidu.com"})
		fmt.Println(v.ValidateOrigin("http://www.baidu.com"))  // true
		fmt.Println(v.ValidateOrigin("https://www.baidu.com")) // true
		fmt.Println(v.ValidateOrigin("https://nssurge.com"))   // false
	}

	{
		v := NewValidator([]string{"https://*.baidu.com"})
		fmt.Println(v.ValidateOrigin("http://www.baidu.com"))  // false
		fmt.Println(v.ValidateOrigin("https://www.baidu.com")) // true
	}
}
