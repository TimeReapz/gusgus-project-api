package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHistory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Order_ID    primitive.ObjectID `bson:"orderId" json:"orderId"`
	STATUS      string             `bson:"status" json:"status"`
	CREATE_DATE time.Time          `bson:"createDate" json:"createDate"`
}
