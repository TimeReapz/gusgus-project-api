package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NAME        string             `json:"name"`
	CREATE_DATE time.Time          `bson:"createDate" json:"createDate"`
	UPDATE_DATE time.Time          `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	ISACTIVE    int                `bson:"isActive,omitempty" json:"isActive,omitempty"`
}

type Users []*User
