package helper

import (
	"errors"
	"regexp"
)

func OnlyLettersValidation(str string) error {
	if !regexp.MustCompile(`^[a-zA-Z ]*$`).MatchString(str) {
		return errors.New("the field must contains only letters")
	}
	return nil
}
