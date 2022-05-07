package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TalentModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Picture   string             `bson:"picture" json:"picture"`
	Job       string             `bson:"job" json:"job"`
	Location  string             `bson:"location" json:"location"`
	LinkedIn  string             `bson:"linkedin" json:"linkedin"`
	GitHub    string             `bson:"github" json:"github"`
	Twitter   string             `bson:"twitter" json:"twitter"`
	Tags      []string           `bson:"tags" json:"tags"`
	Stage     string             `bson:"stage" json:"stage"`
}
