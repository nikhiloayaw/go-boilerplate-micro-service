package routes

import (
	handlerInterfaces "api-gateway/pkg/api/handler/interfaces"
	middlewareInterfaces "api-gateway/pkg/api/middleware/interfaces"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(
	middleware middlewareInterfaces.Middleware,
	authHandler handlerInterfaces.AuthHandler,
	employeeHandler handlerInterfaces.EmployeeHandler,

) http.Handler {

	router := gin.New()

	router.Use(gin.Logger())

	// group with api version
	api := router.Group("/api/v1")

	auth := api.Group("auth")
	{ // for all routes under auth

		RegisterAuthRoutes(auth, authHandler)
	}

	// from here onward all the api should be under the authenticate middleware
	api.Use(middleware.Authenticate("user"))

	// group employee routes
	employees := api.Group("/employees")
	{
		// register employees handler with employees routes
		RegisterEmployeeRoutes(employees, middleware, employeeHandler)
	}

	return router
}
