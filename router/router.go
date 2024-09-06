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
	shortenController *controllers.ShortenController,
	redirectController *controllers.RedirectController,
) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("view/index.html")

	//web
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	//redirect
	r.GET("/redirect/:name", redirectController.Get)

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

		guestRouter := apiRouter.Group("/guest")
		{
			guestRouter.POST("/shorten", shortenController.CreateShortenByPub)
		}

		cmsRouter := apiRouter.Group("/cms")
		{
			cmsRouter.Use(middlewares.DeserializeUser(userRepository))
			cmsRouter.Use(middlewares.AuthorizationAdmin())
			cmsRouter.GET("/shorten", shortenController.GetListShorten)
		}

		userRouter := apiRouter.Group("/user")
		{
			userRouter.Use(middlewares.DeserializeUser(userRepository))
			shortenRouter := userRouter.Group("/shorten")
			{
				shortenRouter.GET("/", shortenController.GetListShorten)
				shortenRouter.POST("/", shortenController.CreateShortenByUser)
				shortenRouter.PUT("/:id", shortenController.UpdateShorten)
			}
		}

	}

	return r
}
