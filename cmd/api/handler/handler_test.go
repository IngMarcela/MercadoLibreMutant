package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidateMutantHandleFindMutant(t *testing.T) {
	dna := InputEvent{[]string{"ATGCGA", "CCGTCC", "TAGAGA", "AAAAGG", "CTTAGA", "TCAGTC"}}
	data, _ := json.Marshal(dna)
	mockValidateMutantUC := &ValidateMutantUCMock{}
	vm := NewHandler(mockValidateMutantUC)
	information, _ := vm.getBodyOfRequest(events.APIGatewayProxyRequest{Body: string(data)})
	mockValidateMutantUC.On("Handler", information).Return(true, nil)

	response, _ := vm.Handle(events.APIGatewayProxyRequest{Body: string(data)})

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestValidateMutantHandleNotFindMutant(t *testing.T) {
	dna := InputEvent{[]string{"ATGCGA", "CCGTCC", "TAGAGA", "AAGAGG", "CTTAGA", "TCAGTC"}}
	data, _ := json.Marshal(dna)
	mockValidateMutantUC := &ValidateMutantUCMock{}
	vm := NewHandler(mockValidateMutantUC)
	information, _ := vm.getBodyOfRequest(events.APIGatewayProxyRequest{Body: string(data)})
	mockValidateMutantUC.On("Handler", information).Return(false, nil)

	response, _ := vm.Handle(events.APIGatewayProxyRequest{Body: string(data)})

	assert.Equal(t, http.StatusForbidden, response.StatusCode)
}

func TestValidateMutantHandleNotFindsMutant(t *testing.T) {
	errorExpected := errors.New("***")
	dna := InputEvent{[]string{"ATGCGA", "CCDTCC", "TAGAGA", "AAGAGG", "CTTAGA", "TCAGTC"}}
	data, _ := json.Marshal(dna)
	mockValidateMutantUC := &ValidateMutantUCMock{}
	vm := NewHandler(mockValidateMutantUC)
	information, _ := vm.getBodyOfRequest(events.APIGatewayProxyRequest{Body: string(data)})
	mockValidateMutantUC.On("Handler", information).Return(false, errorExpected)

	response, _ := vm.Handle(events.APIGatewayProxyRequest{Body: string(data)})

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, errorExpected.Error(), response.Body)
}

func TestValidateMutantHandleNotFindsMutantWithError(t *testing.T) {
	errorExpected := errors.New(MessageWithoutData)
	mockValidateMutantUC := &ValidateMutantUCMock{}
	vm := NewHandler(mockValidateMutantUC)

	response, _ := vm.Handle(events.APIGatewayProxyRequest{})

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, errorExpected.Error(), response.Body)
}
