package mysql

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName  string   `json:"username" gorm:"column:username;unique;type:varchar(100);not null"`
	Password  string   `json:"password,omitempty" gorm:"column:password;type:varchar(100);not null"`
	Name      string   `json:"name" gorm:"column:name;unique;type:varchar(100);not null"`
	Telephone string   `json:"telephone" gorm:"column:telephone;unique;type:varchar(100);not null"`
	IsAdmin   bool     `json:"is_admin" gorm:"column:is_admin;not null;"`
	Tenants   []Tenant `json:"tenants" gorm:"many2many:tenant_user"`
	Groups    []Group  `gorm:"many2many:group_user"`
}

type Tenant struct {
	//ID   primitive.ObjectID `bson:"_id,omitempty"`
	gorm.Model
	Name string `json:"name" gorm:"column:name;unique;type:varchar(100);not null"`
	Code string `json:"code" gorm:"column:code;unique;type:varchar(100);not null"`
	//Admin []primitive.ObjectID `json:"admin"`
	Users []User `json:"-" gorm:"many2many:tenant_user"`
}

type Group struct {
	gorm.Model
	// 关联到租户
	TenantID uint
	Tenant   Tenant
	//	权限组定义
	Code  string `gorm:"column:code;type:varchar(100);not null"`
	Name  string `gorm:"column:name;type:varchar(100);not null"`
	Users []User `gorm:"many2many:group_user"`
}

type Active struct {
	gorm.Model
	Method string `json:"method" gorm:"column:method;type:varchar(10);not null"`
	Path   string `json:"path" gorm:"column:path;type:varchar(200);not null"`
}

type UserToSub struct {
	UserID primitive.ObjectID `bson:"user_id"`
	SubID  primitive.ObjectID `bson:"sub_id"`
}

type CasbinRule struct {
	ID     int    `gorm:"column:id" json:"id"`
	PType  string `gorm:"column:ptype" json:"p_type"`
	Group  string `gorm:"column:v0" json:"group"`
	Type   string `gorm:"column:v1" json:"type"`
	Path   string `gorm:"column:v2" json:"path"`
	Tenant string `gorm:"column:v3" json:"tenant"`
}

func (c *CasbinRule) TableName() string {
	return "casbin_rule"
}
