package di

import (
	"MeliMutant/cmd/api/handler"
	"MeliMutant/pkg"
	"MeliMutant/pkg/validators"
	"MeliMutant/repository"
)

func Initialize() (*handler.Handler, error) {
	sessionProvider := AWSSessionProvider()
	snsProvider := AWSSNSProvider(sessionProvider)
	ValidateChar := validators.NewValidateChar()
	notification := repository.NewNotificationRepository(snsProvider)
	validateMutantUC := pkg.NewValidateMutantUC(ValidateChar, notification)
	handlerMutants := handler.NewHandler(validateMutantUC)
	return handlerMutants, nil
}
