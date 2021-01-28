package validate

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func ValidatorInit() {
	// 初始化校验
	Validate = validator.New()

	Validate.RegisterStructValidation(FooStructLevelValidation, Foo{})
}
