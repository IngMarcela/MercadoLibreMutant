package response

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func response(stringResponse string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: statusCode}
}
func Response200() events.APIGatewayProxyResponse {
	return response("", http.StatusOK)
}
func Response403() events.APIGatewayProxyResponse {
	return response("", http.StatusForbidden)
}
func Response400(stringResponse string) events.APIGatewayProxyResponse {
	return response(stringResponse, http.StatusBadRequest)
}
