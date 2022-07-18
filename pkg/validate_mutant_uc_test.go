package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"MeliMutant/pkg/validators"
)

func TestValidateV1MutantRepositoryFindWithHorizontal(t *testing.T) {
	information := []string{"ATGCGA", "CCGTCC", "TAAAGA", "AAAAGG", "CTTAGA", "TCAGTA"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, _ := vm.Handler(information)

	assert.Equal(t, true, response)
}

func TestValidateV1MutantRepositoryFindWithVertical(t *testing.T) {
	information := []string{"ATGCGA", "CCGTCC", "TAGAGA", "AAGAGG", "CTTAGA", "TCAGTA"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, _ := vm.Handler(information)

	assert.Equal(t, true, response)
}

func TestValidateV1MutantRepositoryNotFind(t *testing.T) {
	information := []string{"ATGCGA", "CCGTCC", "TAAAGA", "AACAGG", "CTTAGA", "TCAGTA"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, _ := vm.Handler(information)

	assert.Equal(t, false, response)
}

func TestValidateV1MutantRepositoryFindWithFirstOblique(t *testing.T) {
	information := []string{"ATGCGA", "CAGTCC", "TAAAGA", "AACAGG", "CTTAGA", "TCAGTG"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, _ := vm.Handler(information)

	assert.Equal(t, true, response)
}

func TestValidateV1MutantRepositoryFindWithSecondOblique(t *testing.T) {
	information := []string{"ATGCGC", "CCGTCC", "TAACGA", "AACAGG", "CTTAGA", "TCAGTG"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, _ := vm.Handler(information)

	assert.Equal(t, true, response)
}

func TestInformationNotAllowedForMutantReading(t *testing.T) {
	errorExpected := messageValidateCaracters
	information := []string{"ADGCGA", "CCGTCC", "TAAAGA", "AAAAGG", "CTTAGA", "TCAGTA"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, err := vm.Handler(information)

	assert.Equal(t, false, response)
	assert.Equal(t, errorExpected, err.Error())
}

func TestIsNotMatrizForReading(t *testing.T) {
	errorExpected := messageValidateDimensionsMatriz
	information := []string{"ATGCA", "CCGTA", "TAAAG", "AAAAT", "CTTAC", "TCAGG"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, err := vm.Handler(information)

	assert.Equal(t, false, response)
	assert.Equal(t, errorExpected, err.Error())
}

func TestIsNotAnArrayWithTheMinimumLengthForReading(t *testing.T) {
	errorExpected := messageValidateMinimumSize
	information := []string{"ATG", "CCG", "TAA"}
	validateDna := validators.NewValidateChar()
	vm := NewValidateMutantUC(validateDna)

	response, err := vm.Handler(information)

	assert.Equal(t, false, response)
	assert.Equal(t, errorExpected, err.Error())
}
