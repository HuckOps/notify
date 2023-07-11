package rbac

import (
	"github.com/HuckOps/notify/src/model/mysql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PK       string `json:"pk"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	gorm.Model
	UserName  string `json:"username" gorm:"column:username"`
	Name      string `json:"name"`
	Telephone string `json:"telephone" gorm:"column:telephone"`
}

type Active struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Method string             `json:"method"`
	Path   string             `json:"path"`
}

type Tenant struct {
	mysql.Tenant
	UserTotal  int `json:"user_total"`
	TokenTotal int `json:"token_total"`
}
