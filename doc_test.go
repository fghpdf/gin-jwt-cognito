package jwt_test

import (
	ginjwt "github.com/akhettar/gin-jwt-cognito"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExampleAuthMiddleware() {

	// Creates a gin router with default middleware:
	router := gin.Default()

	// Create Cognito JWT auth middleware and set it  in all authenticated endpoints
	mw, err := ginjwt.AuthJWTMiddleware("<some_userpool_id>", "region")
	if err != nil {
		panic(err)
	}

	router.GET("/someGet", mw.MiddlewareFunc(), func(c *gin.Context) {
		token := c.MustGet(ginjwt.ContextToken)
		claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
		user := make([]string, 0)
		if email, ok := claims["email"]; ok {
			user = append(user, email.(string))
		}

		if username, ok := claims["username"]; ok {
			user = append(user, username.(string))
		}

		c.JSON(200, gin.H{"user": user})
	})
	router.POST("/somePost", mw.MiddlewareFunc(), func(c *gin.Context) {
		// some implementation
	})
	router.PUT("/somePut", mw.MiddlewareFunc(), func(c *gin.Context) {
		// some implementation
	})

	// By default, it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
