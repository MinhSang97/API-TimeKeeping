//package router
//
//import (
//	"app/handler"
//	appMiddleware "app/middleware"
//	"github.com/labstack/echo/v4"
//)
//
//type API struct {
//	Echo         *echo.Echo
//	AdminHandler handler.AdminHandler
//	RepoHandler  handler.RepoHandler
//}
//
//func (api *API) SetupRouter() {
//	// user
//
//	api.Echo.POST("/user/sign-in", api.AdminHandler.HandleSignIn, appMiddleware.ISAdmin())
//	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
//}

package router

import (
	"app/dbutil"
	"app/handler/admin"
	usersHandler "app/handler/users"
	"app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT | PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delete item by id)
	//viper.SetConfigFile("config.yaml")
	//if err := viper.ReadInConfig(); err != nil {
	//	panic(err)
	//}

	// Your existing code to set up routes and database

	// Register your route, passing the configuration

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	//r.Use(middleware.BasicAuthMiddleware())

	r.GET("/secure", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a secure route"})
	})

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			//admin
			items.POST("/admin/sign-up", handler.AdminSignUp())
			items.POST("/admin/sign-in", handler.AdminSignIn())

			//user
			items.POST("/users/sign-up", middleware.JWTMiddleware(), usersHandler.UsersSignUp())
			items.POST("/users/sign-in", usersHandler.UsersSignIn())
			//items.GET("", handler.GetAllStudent(db))
			//items.GET("/:id", handler.GetId(db))
			//items.PATCH("/:id", handler.Update_One(db))
			//items.DELETE("/:id", handler.Delete_One(db))
			//items.GET("/search", handler.SearchStudents(db))

		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
