package models

import (
	"context"
	"gin-easy/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var UserColl *mongo.Collection

func init() {
	// 连接数据库
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Conf.DB.URI))
	if err != nil {
		log.Fatal(err)
	}
	// 测试连通性
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// 打开/创建数据库
	db := client.Database(config.Conf.DB.Name)
	// 创建集合
	UserColl = db.Collection("user")
}
