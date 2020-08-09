package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"createdAt"`
}
