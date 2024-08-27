package router

import (
	"api-curut-in/controllers"
	"api-curut-in/middlewares"
	"api-curut-in/repository"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	userRepository repository.UsersRepository,
	userController *controllers.UserController,
	authenticationController *controllers.AuthenticationController,
) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("view/index.html")

	//web
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	//api
	r.Use(middlewares.CORSMiddleware())
	apiRouter := r.Group("/api")
	{
		authRouter := apiRouter.Group("/auth")
		{
			authRouter.GET("/users", middlewares.DeserializeUser(userRepository), userController.GetUserByToken)
			authRouter.POST("/login", authenticationController.Login)
			authRouter.POST("/register", authenticationController.Register)
			// authRouter.GET("/users", authenticationController.CheckToken)
			// authRouter.POST("/logout", authenticationController.Logout)
		}
	
	}

	return r
}
