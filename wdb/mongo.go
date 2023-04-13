package wdb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDb mongodb 连接
func InitMongoDb(addr string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(addr)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	fmt.Println("successfully connected and pinged.")
	return client, nil
}

// InsertCollection 写入 db 写入的数据库  collection 写入的文档（表） data 写入的数据
func InsertCollection(m *mongo.Client, db string, collection string, data interface{}) (
	res *mongo.InsertOneResult, err error) {
	c := m.Database(db).Collection(collection)
	return c.InsertOne(context.TODO(), data)
}

// BatchInsertCollection 批量写入
func BatchInsertCollection(m *mongo.Client, db string, collection string, data []interface{}) (
	res *mongo.InsertManyResult, err error) {
	c := m.Database(db).Collection(collection)
	return c.InsertMany(context.TODO(), data)
}

// UpdateOneRecord 修改单条记录
func UpdateOneRecord(m *mongo.Client, db string, collection string, id string, data bson.D) (
	res *mongo.UpdateResult, err error) {
	c := m.Database(db).Collection(collection)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objId}}
	d := bson.D{{"$set", data}}
	res, err = c.UpdateOne(context.TODO(), filter, d)
	return
}

// CountCollection 统计
func CountCollection(m *mongo.Client, db string, collection string, filter map[string]interface{}) (int64, int64, error) {
	c := m.Database(db).Collection(collection)

	esCount, err := c.EstimatedDocumentCount(context.TODO())
	if err != nil {
		return -1, -1, err
	}

	if len(filter) == 0 {
		count, err := c.CountDocuments(context.TODO(), bson.D{{}})
		if err != nil {
			return -1, -1, err
		}
		return esCount, count, nil
	}

	if len(filter) > 0 {
		for k, v := range filter {
			count, err := c.CountDocuments(context.TODO(), bson.D{{k, v}})
			if err != nil {
				return -1, -1, err
			}
			return esCount, count, nil
		}
	}

	return -1, -1, errors.New("查询条件异常")
}

// DeleteOneRecord 删除单条记录
func DeleteOneRecord(m *mongo.Client, db string, collection string, filter map[string]interface{}) (
	res *mongo.DeleteResult, err error) {
	c := m.Database(db).Collection(collection)
	for k, v := range filter {
		filter := bson.D{{k, v}}
		return c.DeleteOne(context.TODO(), filter)
	}
	return
}
