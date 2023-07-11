package rbac

import (
	"github.com/HuckOps/notify/src/db/mysql"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var Enforce *casbin.Enforcer

type CasbinModel struct {
	ID       int
	Ptype    string `json:"ptype" bson:"ptype"`
	RoleName string `json:"rolename" bson:"v0"`
	Path     string `json:"path" bson:"v1"`
	Method   string `json:"method" bson:"v2"`
}

func Init() {
	a, err := gormadapter.NewAdapterByDB(mysql.MySQL.DB())
	if err != nil {
		panic(err)
	}
	//a.LoadPolicy()
	e, err := casbin.NewEnforcer("./conf/rbac_module.conf", a)
	if err != nil {
		panic(err)
	}
	e.LoadPolicy()
	//e.AddPolicy("test", "GET", "/admin/rbac/user", "test")
	e.EnableLog(true)
	e.SavePolicy()
	Enforce = e
}
