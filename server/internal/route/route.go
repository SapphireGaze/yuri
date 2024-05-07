package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sapphiregaze/yuri-server/internal/redirect"
)

type Router struct {
	router *gin.Engine
}

func InitRoutes() Router {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       12 * time.Hour,
	}))

	r := Router{
		router: router,
	}

	r.router.GET("/:redirect", redirect.RedirectHandler)
	r.router.POST("/", redirect.CreateRedirectHandler)

	return r
}

func (r Router) Run(addr ...string) error {
	return r.router.Run()
}
