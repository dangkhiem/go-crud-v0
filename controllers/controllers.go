package controllers

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	// ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	// Password string `bson:"password" json:"password"`
	// CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	// UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func ListUsers(w *gin.ResponseWriter, r *gin.Context) {
	w
}

// func CreateUser() int {
// 	return 2
// }

// func FindUser() int {
// 	return 3
// }

// func UpdateUser() int {
// 	return 4
// }

// func DeleteUser() int {
// 	return 5
// }
