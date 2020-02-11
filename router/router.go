package router

import (
	"net/http"

	"myapi/router/middleware"
	"myapi/service"

	"github.com/gin-gonic/gin"
)

//InitRouter
func InitRouter(g *gin.Engine) {
	var mws []gin.HandlerFunc
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(CORSMiddleware())
	g.Use(mws...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	router := g.Group("/user")
	{
		router.POST("/addUser", service.AddUser)                    //添加用户
		router.GET("/selectUser", service.SelectUser)          //查询用户
		router.GET("/index", service.Index)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
			"Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.String(200, "ok")
			return
		}
		c.Next()
	}
}
