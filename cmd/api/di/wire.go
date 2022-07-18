package di

import (
	"MeliMutant/cmd/api/handler"
	"MeliMutant/pkg"
	"MeliMutant/pkg/validators"
)

func Initialize() (*handler.Handler, error) {
	ValidateChar := validators.NewValidateChar()
	validateMutantUC := pkg.NewValidateMutantUC(ValidateChar)
	handlerMutants := handler.NewHandler(validateMutantUC)
	return handlerMutants, nil
}
