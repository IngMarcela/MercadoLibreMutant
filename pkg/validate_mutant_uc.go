package pkg

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

const messageValidateCaracters = "La información ingresada, No se puede procesar, los Datos no son los permitidos!"
const messageValidateDimensionsMatrix = "la información ingresada, No se puede procesar, no es una matriz de NxN!"
const messageValidateMinimumSize = "la información ingresada, No se puede procesar, La dimension de la matriz debe ser superior a 4 columnas y 4 filas!"

type InputEvent struct {
	Dna []string `json:"dna"`
}

type Adn struct {
	Id string
}

type Message struct {
	Dna    string `json:"dna"`
	Result bool   `json:"result"`
}

type RequestValidationInterface interface {
	IsValidateChar(char string) bool
}

type NotificationInterface interface {
	SendNotification(mutant string) error
}

type ValidateMutantUC struct {
	RequestValidationInterface RequestValidationInterface
	Notification               NotificationInterface
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

	message := Message{
		Dna:    strings.Join(request, " "),
		Result: isMutantResult,
	}

	messageJson, err := json.Marshal(message)

	if err != nil {
		log.Println(err.Error())

		return false, err
	}

	vm.Notification.SendNotification(string(messageJson))

	return isMutantResult, nil
}

func (vm *ValidateMutantUC) validateAndTransformToArray(dna []string) ([][]string, error) {
	matrixDna := make([][]string, len(dna[0]))
	rows := len(dna)
	columns := len(dna[0])
	if rows != columns {
		return matrixDna, errors.New(messageValidateDimensionsMatrix)
	}
	for i := 0; i < len(dna[0]); i++ {
		chars := []rune(dna[i])
		matrixDna[i] = make([]string, len(chars))
		for j := 0; j < len(chars); j++ {
			char := string(chars[j])
			if vm.RequestValidationInterface.IsValidateChar(char) {
				matrixDna[i][j] = char
			} else {
				return matrixDna, errors.New(messageValidateCaracters)
			}
		}
	}

	if rows < 4 && columns < 4 {
		return matrixDna, errors.New(messageValidateMinimumSize)
	}

	return matrixDna, nil
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
	Notification NotificationInterface,
) *ValidateMutantUC {
	return &ValidateMutantUC{
		RequestValidationInterface: RequestValidationInterface,
		Notification:               Notification,
	}
}
