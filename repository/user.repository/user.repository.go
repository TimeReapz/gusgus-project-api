package user_repository

import (
	"context"
	"time"

	"github.com/gusgus-project/database"
	m "github.com/gusgus-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("user")
var ctx = context.Background()

func Insert(user m.User) (m.User, error) {
	user.CREATE_DATE = time.Now()
	user.ISACTIVE = 1

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return Get(id)
}

func Get(id string) (m.User, error) {
	user := m.User{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetJoin(id string) ([]bson.M, error) {
	var users []bson.M

	_id, _ := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return user, err
	// }

	cur, err := collection.Find(ctx, bson.M{"_id": _id})

	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {

		var user bson.M
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func Search(filter interface{}) (m.Users, error) {
	var users m.Users

	if filter == nil {
		filter = bson.M{
			"isActive": 1,
		}
	}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var user m.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func Update(userId string, user m.User) error {
	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":       user.NAME,
			"updateDate": time.Now(),
			"isActive":   1,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {
	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"updateDate": time.Now(),
			"isActive":   0,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}
