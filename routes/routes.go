package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hedagaurav/blog-api/controllers"
)

// SetupRoutes initializes the routes for the application
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// User routes
		userGroup := api.Group("/v1/user")
		{
			userGroup.POST("/register", controllers.Register)
			userGroup.POST("/login", controllers.Login)
		}

		// Post routes
		postGroup := api.Group("/v1/post")
		{
			postGroup.POST("/", controllers.CreatePost)
			postGroup.GET("posts", controllers.GetPosts)
			postGroup.GET("/:id", controllers.GetPost)
			postGroup.PUT("/:id", controllers.UpdatePost)
			postGroup.DELETE("/:id", controllers.DeletePost)

			// todo: make these routes in the v2 version of the API.
			//postGroup.GET("/user/:id", controllers.GetPostsByUserID)
			//postGroup.GET("/search", controllers.SearchPosts)
			//postGroup.GET("/recent", controllers.GetRecentPosts)
			//postGroup.GET("/drafts", controllers.GetDraftPosts)
		}

		// post comment routes
		//commentGroup := api.Group("/api/v1/comment")
		//{
		//	commentGroup.POST("/:post_id", controllers.PostComment)
		//}
	}

	return r
}
