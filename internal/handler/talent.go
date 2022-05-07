package handler

import (
	"crewtest/internal/model"
	"crewtest/internal/repo"
	"crewtest/pkg/response"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
)

type TalentHandler struct {
	talentRepo repo.TalentRepoInterface
}

func NewHandler(repo repo.TalentRepoInterface) *TalentHandler {
	return &TalentHandler{talentRepo: repo}
}

func (h *TalentHandler) List(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var page = "1"
	var limit = "25"

	if len(request.QueryStringParameters["page"]) > 0 {
		page = request.QueryStringParameters["page"]
	}
	if len(request.QueryStringParameters["limit"]) > 0 {
		limit = request.QueryStringParameters["limit"]
	}

	ipage, err := strconv.Atoi(page)
	if err != nil {
		return response.BadRequest(err), err
	}
	ilimit, err := strconv.Atoi(limit)
	if err != nil {
		return response.BadRequest(err), err
	}

	results, err := h.talentRepo.List(ipage, ilimit)
	if err != nil {
		log.Printf("error listing talent documents: %v\n", err)
		return response.Failure(), err
	}

	return response.Success(results), nil
}

func (h *TalentHandler) Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var doc = model.TalentModel{ID: primitive.NewObjectID()}

	err := json.Unmarshal([]byte(request.Body), &doc)
	if err != nil {
		return response.BadRequest(err), err
	}
	err = h.talentRepo.Create(&doc)
	if err != nil {
		log.Printf("error creating talent document: %v\n", err)
		return response.Failure(), err
	}

	return response.Created(map[string]string{"created": "success"}), nil
}
