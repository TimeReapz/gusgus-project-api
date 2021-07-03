package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User_ID     primitive.ObjectID `bson:"userId" json:"userId"`
	OrderItems  []OrderItem        `bson:"orderItems" json:"orderItems"`
	TOTALPRICE  int                `bson:"totalPrice,omitempty" json:"totalPrice,omitempty"`
	DELIVERY    string             `bson:"delivery,omitempty" json:"delivery,omitempty"`
	SCHEDULE    string             `bson:"schedule,omitempty" json:"schedule,omitempty"`
	REMARK      string             `bson:"remark,omitempty" json:"remark,omitempty"`
	CREATE_DATE time.Time          `bson:"createDate" json:"createDate"`
	UPDATE_DATE time.Time          `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	ISACTIVE    int                `bson:"isActive,omitempty" json:"isActive,omitempty"`
}

type OrderItem struct {
	ProdcutId primitive.ObjectID `bson:"productId,omitempty" json:"productId,omitempty"`
	QTY       int                `bson:"qty,omitempty" json:"qty,omitempty"`
}

type Orders []*Order
