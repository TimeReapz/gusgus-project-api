package product_repository

import (
	"context"
	"time"

	"github.com/gusgus-project/database"
	m "github.com/gusgus-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("product")
var ctx = context.Background()

// func UploadFile(file, filename string) {
// 	collection.Gridfs.UploadFile(file, filename)

// 	bucket, err := gridfs.NewBucket(
// 		conn.Database("myfiles"),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	uploadStream, err := bucket.OpenUploadStream(
// 		filename,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer uploadStream.Close()

// 	fileSize, err := uploadStream.Write(data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Write file to DB was successful. File size: %d\n", fileSize)
// }

func Insert(product m.Product) (m.Product, error) {
	product.CREATE_DATE = time.Now()
	product.ISACTIVE = 1

	res, err := collection.InsertOne(ctx, product)
	if err != nil {
		return product, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return Get(id)
}

func Get(id string) (m.Product, error) {
	product := m.Product{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func Search(filter interface{}) (m.Products, error) {
	var products m.Products

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var product m.Product
		err = cur.Decode(&product)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func Update(userId string, product m.Product) error {
	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":       product.NAME,
			"subType":    product.SUBTYPE,
			"price":      product.PRICE,
			"thumbnail":  product.THUMBNAIL,
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
