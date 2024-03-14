package dto

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type (
	AuthRequestDTO struct {
		Password string `json:"password" valid:"required,length(8|255)"`
		Username string `json:"username" valid:"required,matches(^[a-zA-Z0-9_]+$)"`
	}
	AuthResponseDTO struct {
		SessionID string `json:"session_id"`
	}
)

func (authReqDTO *AuthRequestDTO) Validate() []string {
	_, err := govalidator.ValidateStruct(authReqDTO)
	return collectErrors(err)
}

func collectErrors(err error) []string {
	validationErrors := make([]string, 0)
	if err == nil {
		return validationErrors
	}
	var allErrs govalidator.Errors
	if errors.As(err, &allErrs) {
		for _, fld := range allErrs {
			validationErrors = append(validationErrors, fld.Error())
		}
	}
	return validationErrors
}
