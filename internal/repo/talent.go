package repo

import (
	"context"
	"crewtest/internal/model"
	"crewtest/pkg/mongoClient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "talents"
)

//TalentRepoInterface holds contract for talent storage management
type TalentRepoInterface interface {
	List(page, limit int) ([]model.TalentModel, error)
	Create(obj *model.TalentModel) error
}

//TalentRepo represents talent repository
type TalentRepo struct {
	db mongoClient.MongoClientInterface
}

//NewRepo returns TalentRepoInterface object
func NewRepo(db mongoClient.MongoClientInterface) TalentRepoInterface {
	return &TalentRepo{db: db}
}

func (r *TalentRepo) collection() *mongo.Collection {
	return r.db.Storage(collection)
}

//List function returns list of talents
func (r *TalentRepo) List(page, limit int) ([]model.TalentModel, error) {
	opts := options.Find()
	if limit > 100 {
		limit = 100
	}
	if page < 1 {
		page = 1
	}
	page -= 1

	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(page * limit))

	cursor, err := r.collection().Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var results = make([]model.TalentModel, 0)
	for cursor.Next(context.TODO()) {
		result := model.TalentModel{}
		err = cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

//Create function inserts talent document into storage
func (r *TalentRepo) Create(obj *model.TalentModel) error {
	_, err := r.collection().InsertOne(context.Background(), obj)
	return err
}
