package response

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func JSON(data interface{}, statusCode int) events.APIGatewayProxyResponse {
	encoded, _ := json.Marshal(data)

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(encoded),
	}
}

func Created(data interface{}) events.APIGatewayProxyResponse {
	return JSON(data, http.StatusCreated)
}

func Success(data interface{}) events.APIGatewayProxyResponse {
	return JSON(data, http.StatusOK)
}

func BadRequest(err error) events.APIGatewayProxyResponse {
	return JSON(map[string]string{"error": err.Error()}, http.StatusBadRequest)
}

func Failure() events.APIGatewayProxyResponse {
	return JSON(map[string]string{"error": "internal server error"}, http.StatusInternalServerError)
}

func MethodNotAllowed() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
}
