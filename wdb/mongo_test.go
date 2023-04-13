package wdb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestInitMongoDb(t *testing.T) {
	addr := "mongodb://test:123456@127.0.0.1:27017/wailrocket"
	session, err := InitMongoDb(addr)
	defer session.Disconnect(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(session)
}

func TestInsertCollection(t *testing.T) {
	addr := "mongodb://test:123456@127.0.0.1:27017/wailrocket"
	session, err := InitMongoDb(addr)
	defer session.Disconnect(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	type Article struct {
		ID      int
		Title   string
		Content string
	}
	a := &Article{
		ID:      1,
		Title:   "测试",
		Content: "测试测试测试测试测试测试测试",
	}

	res, err := InsertCollection(session, "wailrocket", "admin", a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.InsertedID)
}

func TestBatchInsertCollection(t *testing.T) {
	addr := "mongodb://test:123456@127.0.0.1:27017/wailrocket"
	session, err := InitMongoDb(addr)
	defer session.Disconnect(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	type Article struct {
		ID      int
		Title   string
		Content string
	}

	a := []interface{}{
		Article{
			ID:      2,
			Title:   "测试222",
			Content: "测试测试测试测试测试测试测试",
		},
		Article{
			ID:      3,
			Title:   "测试333",
			Content: "测试测试测试测试测试测试测试",
		},
	}

	res, err := BatchInsertCollection(session, "wailrocket", "admin", a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.InsertedIDs)
}

func TestUpdateOneRecord(t *testing.T) {
	addr := "mongodb://test:123456@127.0.0.1:27017/wailrocket"
	session, err := InitMongoDb(addr)
	defer session.Disconnect(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	type Article struct {
		ID      int
		Title   string
		Content string
	}

	b := bson.D{{"title", "测试333修改"}, {"content", "测试测试测试测试测试测试测试 - 修改修改修改修改修改"}}
	res, err := UpdateOneRecord(session, "wailrocket", "admin", "6437e7649eb92c74e1306392", b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("UpsertedID: %d\n", res.UpsertedID)
	fmt.Printf("ModifiedCount: %d\n", res.ModifiedCount)
}

func TestCountCollection(t *testing.T) {
	addr := "mongodb://test:123456@127.0.0.1:27017/wailrocket"
	session, err := InitMongoDb(addr)
	defer session.Disconnect(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	m := make(map[string]interface{})
	m["title"] = "测试"
	m["content"] = "测试测试测试测试测试测试测试"

	estCount, count, err := CountCollection(session, "wailrocket", "admin", m)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("estCount: %d\n", estCount)
	fmt.Printf("estCount: %d\n", count)
}

func TestDeleteOneRecord(t *testing.T) {
	addr := "mongodb://test:123456@127.0.0.1:27017/wailrocket"
	session, err := InitMongoDb(addr)
	defer session.Disconnect(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	m := make(map[string]interface{})
	m["title"] = "测试"
	m["content"] = "测试测试测试测试测试测试测试"
	res, err := DeleteOneRecord(session, "wailrocket", "admin", m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.DeletedCount)
}
