package main

import (
	"authena/models/user"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

//4 REST routes:
//
// 1 - returns a pair of Access-Refresh tokens for the user
//	   with GUID specified in request
//
// 2 - executes Refresh operation on a pair of Access-Refresh
//     token
//
// 3 - deletes Refresh token from the db
//
// 4 - deletes all Refresh tokens from the db for a specified user
//
//
var router = gin.Default()

func main() {
	router.POST("/login", login)
	log.Fatal(router.Run(os.Getenv("PORT")))
}

//func login

var user = &User
