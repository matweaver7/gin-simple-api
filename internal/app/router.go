package app

import (
	"github.com/gin-gonic/gin"
	"matweaver.com/simple-rest-api/config"
	"matweaver.com/simple-rest-api/internal/controllers"
)

type Router interface {
	AddRoutes()
}

type RouterConfig struct {
	engine *gin.Engine
	controllers *controllers.Controllers
}

type router struct {
	engine	*gin.Engine
	controllers *controllers.Controllers
}

func NewRouter(cfg *RouterConfig) Router {
	return &router {
		controllers: cfg.controllers,
		engine: cfg.engine,	
	}
}


func (r *router) AddRoutes() {
	baseGroup := r.engine.Group("/rest") {
		baseGroup.GET("/ping", func() => {fmt.Println("cool")})
	}
} 

