package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileUpload struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PATH        string             `json:"path"`
	NAME        string             `json:"name"`
	TYPE        string             `json:"type"`
	SIZE        int                `json:"size"`
	CREATE_DATE time.Time          `bson:"createDate" json:"createDate"`
	UPDATE_DATE time.Time          `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	ISACTIVE    int                `bson:"isActive,omitempty" json:"isActive,omitempty"`
}

type FileUploads []*FileUpload
