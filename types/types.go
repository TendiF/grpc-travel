package types

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	jwt.StandardClaims
	Uid string
}

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt  primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	ModifiedAt primitive.DateTime `bson:"modified_at,omitempty" json:"modified_at"`
	Username   string             `bson:"username,omitempty" json:"username"`
	Phone      string             `bson:"phone,omitempty" json:"phone"`
	Password   string             `bson:"password,omitempty" json:"password"`
}
