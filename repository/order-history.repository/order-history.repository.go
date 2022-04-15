package order_repository

import (
	"context"
	"time"

	"github.com/gusgus-project/database"
	m "github.com/gusgus-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("orderHistory")
var ctx = context.Background()

func Insert(orderHis m.OrderHistory) (m.OrderHistory, error) {
	orderHis.STATUS = "จัดส่งแล้ว"
	orderHis.CREATE_DATE = time.Now()

	res, err := collection.InsertOne(ctx, orderHis)
	if err != nil {
		return orderHis, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return Get(id)
}

func Get(id string) (m.OrderHistory, error) {
	orderHis := m.OrderHistory{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return orderHis, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&orderHis)
	if err != nil {
		return orderHis, err
	}

	return orderHis, nil
}

// func GetToday() (m.OrderHistory, error) {
// 	orderHis := m.OrderHistory{}

// 	_id, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return orderHis, err
// 	}

// 	err = collection.FindOne(ctx, bson.M{"_id": time.Now()}).Decode(&orderHis)
// 	if err != nil {
// 		return orderHis, err
// 	}

// 	return orderHis, nil
// }
