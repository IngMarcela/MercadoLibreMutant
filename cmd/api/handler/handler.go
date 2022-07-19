package handler

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"MeliMutant/cmd/api/response"
	"MeliMutant/pkg/validators"
)

const MessageWithoutData = "no stats was provided in the HTTP body"

type InputEvent struct {
	Dna []string `json:"dna"`
}

type ValidateMutantUCInterface interface {
	Handler(request []string) (bool, error)
}

type Handler struct {
	validateMutantUC ValidateMutantUCInterface
}

func (h *Handler) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if validators.IsInvalidRequest(len(request.Body)) {
		return response.Response400(MessageWithoutData), nil
	}

	requestInformation, err := h.getBodyOfRequest(request)

	if err != nil {
		return response.Response400(err.Error()), nil
	}

	resultMutant, err := h.validateMutantUC.Handler(requestInformation)

	if err != nil {
		return response.Response400(err.Error()), nil
	}

	if resultMutant {
		return response.Response200(), nil
	}

	return response.Response403(), nil
}

func (h *Handler) getBodyOfRequest(request events.APIGatewayProxyRequest) ([]string, error) {
	requestJSON := InputEvent{}
	err := json.Unmarshal([]byte(request.Body), &requestJSON)
	if err != nil {
		log.Println(err.Error())

		return nil, err
	}
	return requestJSON.Dna, nil
}

func NewHandler(
	uc ValidateMutantUCInterface,
) *Handler {
	return &Handler{
		validateMutantUC: uc,
	}
}
