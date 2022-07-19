package pkg

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"MeliMutant/pkg/validators"
)

const functionSendNotification = "SendNotification"

func TestValidateV1MutantRepositoryFindWithHorizontal(t *testing.T) {
	information := []string{"ATGCGA", "CCGTCC", "TAAAGA", "AAAAGG", "CTTAGA", "TCAGTA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, _ := validateMutant.Handler(information)

	assert.Equal(t, true, response)
}

func TestValidateV1MutantRepositoryFindWithVertical(t *testing.T) {
	information := []string{"ATGCGA", "CCGTCC", "TAGAGA", "AAGAGG", "CTTAGA", "TCAGTA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, _ := validateMutant.Handler(information)

	assert.Equal(t, true, response)
}

func TestValidateV1MutantRepositoryNotFind(t *testing.T) {
	information := []string{"ATVCGA", "CCGTCC", "TAGAGA", "AAGAGG", "CTTAGA", "TCAGTA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: false,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, _ := validateMutant.Handler(information)

	assert.Equal(t, false, response)
}

func TestValidateV1MutantRepositoryFindWithFirstOblique(t *testing.T) {
	information := []string{"ATGCGA", "CAGTCC", "TAAAGA", "AAGAGG", "CTTAGA", "TCAGTA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, _ := validateMutant.Handler(information)

	assert.Equal(t, true, response)
}

func TestValidateV1MutantRepositoryFindWithSecondOblique(t *testing.T) {
	information := []string{"ATGCGA", "CCGTAC", "TAGAGA", "AAAGGG", "CTTAGA", "TCAGTA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, _ := validateMutant.Handler(information)

	assert.Equal(t, true, response)
}

func TestInformationNotAllowedForMutantReading(t *testing.T) {
	errorExpected := messageValidateCaracters
	information := []string{"ADGCGA", "CCGTCC", "TAAAGA", "AAAAGG", "CTTAGA", "TCAGTA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, err := validateMutant.Handler(information)

	assert.Equal(t, false, response)
	assert.Equal(t, errorExpected, err.Error())
}

func TestIsNotMatrixForReading(t *testing.T) {
	errorExpected := messageValidateDimensionsMatrix
	information := []string{"ATGCA", "CCGTA", "TAAAG", "AAAAT", "CTTAC", "TCAGG"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, err := validateMutant.Handler(information)

	assert.Equal(t, false, response)
	assert.Equal(t, errorExpected, err.Error())
}

func TestIsNotAnArrayWithTheMinimumLengthForReading(t *testing.T) {
	errorExpected := messageValidateMinimumSize
	information := []string{"ATG", "CCG", "TAA"}
	message := Message{
		Dna:    strings.Join(information, " "),
		Result: true,
	}
	messageJson, _ := json.Marshal(message)
	validateDna := validators.NewValidateChar()
	mockNotification := NotificationMock{}
	mockNotification.On(functionSendNotification, string(messageJson)).Return(nil)
	validateMutant := NewValidateMutantUC(validateDna, &mockNotification)

	response, err := validateMutant.Handler(information)

	assert.Equal(t, false, response)
	assert.Equal(t, errorExpected, err.Error())
}
