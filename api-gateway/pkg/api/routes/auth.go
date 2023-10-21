package routes

import (
	handlerInterface "api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(auth *gin.RouterGroup, authHandler handlerInterface.AuthHandler) {

	signUp := auth.Group("/sign-up")
	{ // all sign up related routes, like sign-up verify, sign up with google etc..
		signUp.POST("", authHandler.SignUp)

	}

	signIn := auth.Group("/sign-in")
	{
		signIn.POST("", authHandler.SignIn)
	}
}
