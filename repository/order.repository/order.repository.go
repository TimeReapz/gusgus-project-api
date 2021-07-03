package order_repository

import (
	"context"
	"time"

	"github.com/gusgus-project/database"
	m "github.com/gusgus-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = database.GetCollection("order")
var ctx = context.Background()

func Insert(order m.Order) (m.Order, error) {
	order.CREATE_DATE = time.Now()
	order.ISACTIVE = 1

	res, err := collection.InsertOne(ctx, order)
	if err != nil {
		return order, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return Get(id)
}

func Get(id string) (m.Order, error) {
	order := m.Order{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return order, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func GetJoin(id string) ([]bson.M, error) {
	var orders []bson.M

	_id, _ := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return order, err
	// }

	cur, err := collection.Find(ctx, bson.M{"_id": _id})

	if err != nil {
		return orders, err
	}

	for cur.Next(ctx) {

		var order bson.M
		err = cur.Decode(&order)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func Search(filter interface{}) ([]bson.M, error) {
	var orders []bson.M

	if filter == nil {
		filter = bson.M{
			"isActive": 1,
		}
	}

	//matchStage := bson.D{{"$match", filter}}
	// lookup := bson.D{
	// 	{"$match": bson.D{{"_id": "56b9df0c1e930a99cb2c33e9"}}},
	// }

	// const agg = bson.D{
	//   {
	// 	  "$match": {
	// 	"_id": primitive.ObjectID("60e02f8c169043bfd00784dc"),
	//   }},
	//   {"$lookup": bson.D{
	// 	  {
	// 	"from": "user",
	// 	"localField": "userId",
	// 	"foreignField": "_id",
	// 	"as": "userModel",
	//   	}
	//   }}, {$unwind: {
	// 	path: '$orderItems',
	// 	preserveNullAndEmptyArrays: true
	//   }}, {$lookup: {
	// 	from: 'product',
	// 	localField: 'orderItems.productId',
	// 	foreignField: '_id',
	// 	as: 'orderItems.productModel'
	//   }}, {$unwind: {
	// 	path: '$orderItems.productModel',
	// 	preserveNullAndEmptyArrays: true
	//   }}, {$group: {
	// 	_id:ObjectId('60e02f8c169043bfd00784dc'),
	// 	"userId": { "$first": "$userId" },
	// 	"userModel": { "$first": "$userModel" },
	// 	orderItems: {
	// 	  '$push': '$orderItems'
	// 	},
	// 	"totalPrice": { "$first": "$totalPrice" },
	// 	"delivery": { "$first": "$delivery" },
	// 	"schedule": { "$first": "$schedule" },
	//   }}, {$unwind: {
	// 	path: '$userModel',
	// 	preserveNullAndEmptyArrays: true
	//   }}}

	pipeline := mongo.Pipeline{
		{{"$lookup", bson.D{{"from", "user"}, {"localField", "userId"}, {"foreignField", "_id"}, {"as", "userModel"}}}},
		{{"$unwind", "$userModel"}},
		{{"$unwind", bson.D{{"path", "$orderItems"}, {"preserveNullAndEmptyArrays", true}}}},
		{{"$lookup", bson.D{{"from", "product"}, {"localField", "orderItems.productId"}, {"foreignField", "_id"}, {"as", "orderItems.productModel"}}}},
		{{"$unwind", bson.D{{"path", "$orderItems.productModel"}, {"preserveNullAndEmptyArrays", true}}}},
		{{"$group", bson.D{{"_id", "$_id"}, {"userId", bson.D{{"$first", "$userId"}}}, {"userModel", bson.D{{"$first", "$userModel"}}}, {"orderItems", bson.D{{"$push", "$orderItems"}}}, {"totalPrice", bson.D{{"$first", "$totalPrice"}}}, {"delivery", bson.D{{"$first", "$delivery"}}}, {"schedule", bson.D{{"$first", "$schedule"}}}}}},
	}

	cur, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var order bson.M
		err = cur.Decode(&order)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func Update(orderId string, order m.Order) error {
	oid, _ := primitive.ObjectIDFromHex(orderId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"userId":     order.User_ID,
			"orderItems": order.OrderItems,
			"totalPrice": order.TOTALPRICE,
			"delivery":   order.DELIVERY,
			"schedule":   order.SCHEDULE,
			"remark":     order.REMARK,
			"updataDate": time.Now(),
			"isActive":   1,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(orderId string) error {
	oid, _ := primitive.ObjectIDFromHex(orderId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"updataDate": time.Now(),
			"isActive":   0,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}
