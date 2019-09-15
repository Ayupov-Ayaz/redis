package services

import(
	"errors"
	"fmt"
	. "gopkg.in/go-playground/validator.v9"
)

type ValidatorService interface {
	ValidateStruct(i interface{}) error
}

type validatorService struct {
	validate *Validate
}

func (srv *validatorService) ValidateStruct(i interface{}) error {
	err := srv.validate.Struct(i)
	if vErr, ok := err.(ValidationErrors); ok {
		return errors.New(parseFieldValidationErrors(vErr))
	}
	return err
}

func NewValidatorService() ValidatorService {
	return &validatorService{
		validate: New(),
	}
}

func parseFieldValidationErrors(err ValidationErrors) string { //todo: дополнить мб
	messages := ""
	for _, e := range err {
		switch e.Tag() {
		case "required":
			messages += fmt.Sprintf("field '%s' most be %s. ", e.Field(), e.Tag())
		case "min", "max":
			messages += fmt.Sprintf("for '%s' %s length %s chars(elements). ", e.Field(), e.Tag(), e.Param())
		case "eqfield":
			messages += fmt.Sprintf("field '%s' must be equal '%s'. ", e.Field(), e.Param())
		default:
			messages += fmt.Sprintf("%s - not created parse validation for this tag.  ", e.Tag())
		}
	}
	return messages
}