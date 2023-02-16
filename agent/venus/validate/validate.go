package validate

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate
var once sync.Once

func init() {
	once.Do(func() {
		Validate = validator.New()
		Validate.SetTagName("binding")
	})
}
