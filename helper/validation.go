package helper

import (
	"errors"
	"fmt"
	"musiclab-be/features/auth"

	"github.com/go-playground/validator/v10"
)

type UserValidate struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=3,alphanum"`
}

func CoreToRegVal(data auth.Core) UserValidate {
	return UserValidate{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
func RegistrationValidate(data auth.Core) error {
	validate := validator.New()
	val := CoreToRegVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := ""
			if e.Field() == "Password" && e.Value() != "" {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
			if e.Value() == "" {
				vlderror = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
				return errors.New(vlderror)
			} else {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
		}
	}
	return nil
}
