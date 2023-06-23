package mysql

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string `json:"username" gorm:"column:username;unique;type:varchar(100);not null"`
	Password  string `json:"password,omitempty" gorm:"column:password;type:varchar(100);not null"`
	Name      string `json:"name" gorm:"column:name;unique;type:varchar(100);not null"`
	Telephone string `json:"telephone" gorm:"column:telephone;unique;type:varchar(100);not null"`
}

type Active struct {
	gorm.Model
	Method string `json:"method" gorm:"column:method;type:varchar(10);not null"`
	Path   string `json:"path" gorm:"column:path;type:varchar(200);not null"`
}

type Tenant struct {
	//ID   primitive.ObjectID `bson:"_id,omitempty"`
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
	//Admin []primitive.ObjectID `json:"admin"`
}

type Sub struct {
	//ID     primitive.ObjectID `bson:"_id,omitempty"`
	gorm.Model
	Name   string             `json:"name"`
	Code   string             `json:"code"`
	Tenant primitive.ObjectID `bson:"tenant"`
}

type UserToSub struct {
	UserID primitive.ObjectID `bson:"user_id"`
	SubID  primitive.ObjectID `bson:"sub_id"`
}
