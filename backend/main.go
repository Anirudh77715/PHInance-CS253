package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sayymeer/feinance-backend/router"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigin := os.Getenv("URL") // set this to your Vercel frontend URL

	// Optional: allow multiple origins separated by comma
	// Example: URL="https://a.vercel.app,https://b.vercel.app"
	allowedOrigins := []string{}
	if allowedOrigin != "" {
		for _, o := range strings.Split(allowedOrigin, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				allowedOrigins = append(allowedOrigins, o)
			}
		}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		// If URL env is empty, you can allow all (not recommended for production)
		// Better: set URL to your actual frontend domain.
		if len(allowedOrigins) == 0 {
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			// Allow only matched origins
			for _, o := range allowedOrigins {
				if origin == o {
					c.Header("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		c.Header("Vary", "Origin")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// If you use cookies/sessions, enable this and DO NOT use "*" origin
		// c.Header("Access-Control-Allow-Credentials", "true")

		// Preflight request
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Uncomment for production logs
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(CORSMiddleware())

	// Routes
	router.FinRoutes(r)

	// IMPORTANT: bind to platform PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Listen on 0.0.0.0 for cloud deployments
	if err := r.Run("0.0.0.0:" + port); err != nil {
		panic(err)
	}
}
