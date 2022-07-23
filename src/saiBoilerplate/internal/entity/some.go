package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Some struct {
	ID  primitive.ObjectID `bson:"_id" json:"id" example:"d234234ewf334"`
	Key string             `bson:"key" json:"key" example:"we423JJ8w"`
}
