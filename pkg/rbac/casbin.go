package rbac

import (
	"fmt"
	"github.com/HuckOps/notify/src/config"
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?retryWrites=true&w=majority",
		config.Config.DB.Mongo.User, config.Config.DB.Mongo.Password, config.Config.DB.Mongo.Host, config.Config.DB.Mongo.Port, config.Config.DB.Mongo.DB)

	mongoClientOption := mongooptions.Client().ApplyURI(mongoURL)
	a, err := mongodbadapter.NewAdapterWithClientOption(mongoClientOption, config.Config.DB.Mongo.DB)
	if err != nil {
		panic(err)
	}
	//a.LoadPolicy()
	e, err := casbin.NewEnforcer("/home/huck/notify/conf/rbac_module.conf", a)
	if err != nil {
		panic(err)
	}
	e.LoadPolicy()
	e.EnableLog(true)
	e.SavePolicy()
}
