package rbac

import (
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
	ID        uint   `json:"id"`
	UserName  string `json:"username" gorm:"column:username"`
	Name      string `json:"name"`
	Telephone string `json:"telephone" gorm:"column:telephone"`
}

type Active struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Method string             `json:"method"`
	Path   string             `json:"path"`
}
