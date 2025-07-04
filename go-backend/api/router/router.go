package router

import (
	"github.com/gin-gonic/gin"

	"go-backend/api/handlers/auth"
	"go-backend/api/handlers/request"
	"go-backend/api/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST(
			"/onboard",
			middlewares.AuthenticateTokenMiddleware(),
			middlewares.AuthoriseAdminMiddleware(),
			auth.Onboard,
		)
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/logout", auth.Logout)
	}

	resourceRoutes := r.Group("/api")
	resourceRoutes.Use(middlewares.AuthenticateTokenMiddleware())
	{
		resourceRoutes.GET(
			"/summary",
			middlewares.AuthoriseAdminMiddleware(),
			request.Summary,
		)
		resourceRoutes.POST(
			"/request",
			request.Request,
		)
		resourceRoutes.GET(
			"/admin/:id",
			middlewares.AuthoriseAdminMiddleware(),
			middlewares.AuthoriseUserMiddleware(),
			request.Admin,
		)
		resourceRoutes.GET(
			"/employee/:id",
			middlewares.AuthoriseUserMiddleware(),
			request.Employee,
		)
		resourceRoutes.PATCH(
			"/request/:id/status",
			middlewares.AuthoriseAdminMiddleware(),
			request.UpdateRequestStatus,
		)
		resourceRoutes.PATCH(
			"/request/:id/remarks",
			middlewares.AuthoriseAdminMiddleware(),
			request.UpdateRequestRemarks,
		)
	}

	return r
}
