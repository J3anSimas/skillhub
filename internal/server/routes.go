package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"io/fs"
	"train/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))
	staticFiles, _ := fs.Sub(web.Files, "assets")
	r.StaticFS("/assets", http.FS(staticFiles))
	r.GET("/login", func(c *gin.Context) {
		templ.Handler(web.LoginForm()).ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/health", s.healthHandler)
	private := r.Group("/")
	private.Use(isAuthenticated())
	{
		private.GET("/", func(c *gin.Context) {
			templ.Handler(web.HelloForm()).ServeHTTP(c.Writer, c.Request)
		})
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	// c.JSON(http.StatusOK, s.db.Health())
	c.JSON(http.StatusOK, "")
}
