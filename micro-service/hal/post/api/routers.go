

package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	userApi "github.com/flexzuu/benchmark/micro-service/rest/user/openapi/client"

	"github.com/flexzuu/benchmark/micro-service/rest/post/repo"
	"github.com/flexzuu/benchmark/micro-service/rest/post/repo/inmemmory"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

//dependencies
var postRepo repo.Post
var userServiceClient *userApi.APIClient

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	postRepo = inmemmory.NewRepo()

	userServiceAddress := os.Getenv("USER_SERVICE")
	if userServiceAddress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	userCfg := userApi.NewConfiguration()
	userCfg.BasePath = fmt.Sprintf("http://%s", userServiceAddress)
	userServiceClient = userApi.NewAPIClient(userCfg)

	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			router.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			router.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"CreatePost",
		strings.ToUpper("Post"),
		"/post",
		CreatePost,
	},

	{
		"DeletePost",
		strings.ToUpper("Delete"),
		"/post/:id",
		DeletePost,
	},

	{
		"GetPostById",
		strings.ToUpper("Get"),
		"/post/:id",
		GetPostById,
	},

	{
		"ListPosts",
		strings.ToUpper("Get"),
		"/post",
		ListPosts,
	},
}
