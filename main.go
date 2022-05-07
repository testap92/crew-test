package main

import (
	"crewtest/config"
	"crewtest/internal/handler"
	"crewtest/internal/repo"
	"crewtest/pkg/mongoClient"
	"crewtest/pkg/response"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cfg := config.NewConfig()
	mongoCli := mongoClient.Connect(cfg.DB.DBName, cfg.GetDBUrl())
	talentRepo := repo.NewRepo(mongoCli)
	talentHandler := handler.NewHandler(talentRepo)

	log.Println("Starting Lambda...")
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		if request.HTTPMethod == http.MethodPost {
			return talentHandler.Create(request)
		}
		if request.HTTPMethod == http.MethodGet {
			return talentHandler.List(request)
		}
		return response.MethodNotAllowed()
	})
}
