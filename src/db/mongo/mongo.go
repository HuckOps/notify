package mongo

import (
	"context"
	"fmt"
	"github.com/HuckOps/notify/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mgo struct {
	context context.Context
	db      *mongo.Database
	//reloadChan chan bool
}

var Mongo *mgo

func init() {
	Mongo = &mgo{context: context.Background()}
}

func (m *mgo) Load() {

	var mongoURL string
	mongoURL = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?retryWrites=true&w=majority",
		config.Config.DB.Mongo.User, config.Config.DB.Mongo.Password, config.Config.DB.Mongo.Host, config.Config.DB.Mongo.Port, config.Config.DB.Mongo.DB)
	fmt.Println(mongoURL)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	m.db = client.Database("notify")

}

func (m *mgo) DB() *mongo.Database {
	m.db.Collection("user")
	return m.db
}

func SearchByPagination(collection *mongo.Collection, skip int, limit int, match bson.M) (*mongo.Cursor, error) {
	pipeline := []bson.M{
		// 匹配条件（根据需求修改）
		bson.M{
			"$match": match,
		},
		// 获取总数和分页结果
		bson.M{
			"$facet": bson.M{
				"totalCount": []bson.M{
					bson.M{"$count": "total"},
				},
				"items": []bson.M{
					// 分页设置
					bson.M{"$skip": skip},
					bson.M{"$limit": limit},
				},
			},
		},
		// 重塑输出格式
		bson.M{
			"$unwind": "$totalCount",
		},
		bson.M{
			"$project": bson.M{
				"total": "$totalCount.total",
				"items": 1,
			},
		},
	}
	return collection.Aggregate(context.TODO(), pipeline)
}
