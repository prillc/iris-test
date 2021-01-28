package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Foo struct {
	Id   int    `json:"id" validate:"gte=1"`
	Desc string `json:"desc"`
}

func FooStructLevelValidation(sl validator.StructLevel) {
	fmt.Println("FooStructLevelValidation")
}
