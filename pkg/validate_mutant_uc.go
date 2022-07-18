package pkg

import (
	"errors"
)

const messageValidateCaracters = "La información ingresada, No se puede procesar, los Datos no son los permitidos!"
const messageValidateDimensionsMatriz = "la información ingresada, No se puede procesar, no es una matriz de NxN!"
const messageValidateMinimumSize = "la información ingresada, No se puede procesar, La dimension de la matriz debe ser superior a 4 columnas y 4 filas!"

type InputEvent struct {
	Dna []string `json:"stats"`
}

type Adn struct {
	Id string
}

type RequestValidationInterface interface {
	IsValidateChar(char string) bool
}

type ValidateMutantUC struct {
	RequestValidationInterface RequestValidationInterface
}

func (vm *ValidateMutantUC) Handler(request []string) (bool, error) {
	mutant, err := vm.execute(request)

	if err != nil {
		return false, err
	}

	return mutant, nil
}

func (vm *ValidateMutantUC) execute(request []string) (bool, error) {
	informationDna, err := vm.validateAndTransformToArray(request)

	if err != nil {
		return false, err
	}

	isMutantResult := isMutant(informationDna)
	if !isMutantResult {
		return false, nil
	}

	return isMutantResult, nil
}

func (vm *ValidateMutantUC) validateAndTransformToArray(dna []string) ([][]string, error) {
	matrizDna := make([][]string, len(dna[0]))
	rows := len(dna)
	columns := len(dna[0])
	if rows != columns {
		return matrizDna, errors.New(messageValidateDimensionsMatriz)
	}
	for i := 0; i < len(dna[0]); i++ {
		chars := []rune(dna[i])
		matrizDna[i] = make([]string, len(chars))
		for j := 0; j < len(chars); j++ {
			char := string(chars[j])
			if vm.RequestValidationInterface.IsValidateChar(char) {
				matrizDna[i][j] = char
			} else {
				return matrizDna, errors.New(messageValidateCaracters)
			}
		}
	}

	if rows < 4 && columns < 4 {
		return matrizDna, errors.New(messageValidateMinimumSize)
	}

	return matrizDna, nil
}

func isMutant(dna [][]string) bool {
	rta := false
	rta = isHorizontal(dna, rta)
	if !rta {
		rta = isVertical(dna, rta)
	}
	if !rta {
		rta = isOblique(dna, rta)
	}

	return rta
}

func isHorizontal(dna [][]string, rta bool) bool {
	for i := 0; i < len(dna[0]); i++ {
		for j := 0; j < len(dna[0])-3; j++ {
			if dna[i][j] == dna[i][j+1] &&
				dna[i][j] == dna[i][j+2] &&
				dna[i][j] == dna[i][j+3] {
				rta = true

			}
		}
	}

	return rta
}

func isVertical(dna [][]string, rta bool) bool {
	for i := 0; i < len(dna[0])-3; i++ {
		for j := 0; j < len(dna[0]); j++ {
			if dna[i][j] == dna[i+1][j] &&
				dna[i][j] == dna[i+2][j] &&
				dna[i][j] == dna[i+3][j] {
				rta = true
			}
		}
	}

	return rta
}

func isOblique(dna [][]string, rta bool) bool {
	for i := 0; i < len(dna[0])-3; i++ {
		for j := 0; j < len(dna[0])-3; j++ {
			if dna[i][j] == dna[i+1][j+1] &&
				dna[i][j] == dna[i+2][j+2] &&
				dna[i][j] == dna[i+3][j+3] {
				rta = true
			}
		}
	}

	for i := 0; i < len(dna[0])-3; i++ {
		for j := len(dna[0]) - 1; j > 2; j-- {

			if dna[i][j] == dna[i+1][j-1] &&
				dna[i][j] == dna[i+2][j-2] &&
				dna[i][j] == dna[i+3][j-3] {
				rta = true
			}
		}
	}

	return rta
}

func NewValidateMutantUC(
	RequestValidationInterface RequestValidationInterface,
) *ValidateMutantUC {
	return &ValidateMutantUC{
		RequestValidationInterface: RequestValidationInterface,
	}
}
