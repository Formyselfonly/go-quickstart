package controller

import (
	// "go-quickstart/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(s *ghttp.Server) {
	// API group with prefix
	api := s.Group("/api/v1")
	
	// User routes
	userController := &UserController{}
	api.POST("/users", userController.Create)
	api.GET("/users/{id}", userController.Get)
	api.GET("/users", userController.List)
	api.PUT("/users/{id}", userController.Update)
	api.DELETE("/users/{id}", userController.Delete)
	
	// Health check endpoint
	api.GET("/health", func(r *ghttp.Request) {
		r.Response.WriteJson(g.Map{
			"status":  "ok",
			"message": "GoFrame Backend API is running",
		})
	})
	
	// API documentation endpoint
	api.GET("/docs", func(r *ghttp.Request) {
		r.Response.WriteJson(g.Map{
			"title":       "GoFrame Backend API",
			"version":     "1.0.0",
			"description": "A comprehensive backend API built with GoFrame",
			"endpoints": g.Map{
				"users": g.Map{
					"POST":   "/api/v1/users - Create a new user",
					"GET":    "/api/v1/users - List users with pagination",
					"GET":    "/api/v1/users/{id} - Get user by ID",
					"PUT":    "/api/v1/users/{id} - Update user",
					"DELETE": "/api/v1/users/{id} - Delete user",
				},
				"health": g.Map{
					"GET": "/api/v1/health - Health check",
				},
			},
		})
	})
}
