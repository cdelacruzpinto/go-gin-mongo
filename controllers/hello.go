package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cdelacruzpinto/types"
	"github.com/gin-gonic/gin"
)
var identityKey = "id"

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
	  "userID":   claims[identityKey],
	  "userName": user.(*types.User).UserName,
	  "text":     "Hello World.",
	})
  }