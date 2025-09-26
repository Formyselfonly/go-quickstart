package main

import (
	"context"
	"go-quickstart/internal/controller"
	"go-quickstart/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	
	// Initialize services
	service.Init(ctx)
	
	// Start HTTP server
	s := g.Server()
	
	// Add middleware
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Use(ghttp.MiddlewareCORS)
	
	// Register routes
	controller.RegisterRoutes(s)
	
	// Set server configuration
	s.SetPort(8000)
	s.SetServerRoot(".")
	
	// Add a simple welcome route
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteJson(g.Map{
			"message": "Welcome to GoFrame Backend API!",
			"version": "1.0.0",
			"docs":    "/api/v1/docs",
			"health":  "/api/v1/health",
		})
	})
	
	// Start server
	s.Run()
}