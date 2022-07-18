package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"MeliMutant/cmd/api/response"
	"MeliMutant/pkg/validators"
)

const MessageWithoutData = "no stats was provided in the HTTP body"

type InputEvent struct {
	Dna []string `json:"stats"`
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

	requestInformation := h.getBodyOfRequest(request)
	resultMutant, err := h.validateMutantUC.Handler(requestInformation)

	if err != nil {
		return response.Response400(err.Error()), nil
	}

	if resultMutant {
		return response.Response200(), nil
	}

	return response.Response403(), nil
}

func (h *Handler) getBodyOfRequest(request events.APIGatewayProxyRequest) []string {
	requestJSON := InputEvent{}
	json.Unmarshal([]byte(request.Body), &requestJSON)
	return requestJSON.Dna
}

func NewHandler(
	uc ValidateMutantUCInterface,
) *Handler {
	return &Handler{
		validateMutantUC: uc,
	}
}
