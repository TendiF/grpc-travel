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
	FirstName  string             `bson:"first_name,omitempty" json:"first_name"`
	LastName   string             `bson:"last_name,omitempty" json:"last_name"`
	Gender     string             `bson:"gender,omitempty" json:"gender"`
	Email      string             `bson:"email,omitempty" json:"email"`
	Password   string             `bson:"password,omitempty" json:"password"`
	Address    string             `bson:"address,omitempty" json:"address"`
}
