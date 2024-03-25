package middlewares

import (
	"github.com/gin-gonic/gin"
)

// The `CORS` middleware function is designed to handle Cross-Origin Resource Sharing (CORS) in our backend services. It allows secure communication between client-side web applications and our server, even if they're hosted on different domains.
func CORS(domain string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", domain)
		ctx.Writer.Header().Set("Access-Control-Max-Age", "0")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, content-encoding, Content-Encoding, x-action-name")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, X-Total-Number")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("X-Frame-Options", "DENY")
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		ctx.Writer.Header().Set("X-Download-Options", "noopen")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}

		ctx.Next()
	}
}
