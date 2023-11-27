package param_verify

import (
	"github.com/go-playground/validator/v10"
)

func NameValid(field validator.FieldLevel) bool {
	// 一个简单的验证器例子，不允许字段值为 admin
	if s := field.Field().String(); s != "" {
		if s == "admin" {
			return false
		}
	}
	return true
}
