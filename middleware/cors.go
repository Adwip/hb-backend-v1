package middleware

import "github.com/gin-gonic/gin"
import "os"

type cors struct {
}

func CORS() *cors {
	corsObj := &cors{}
	return corsObj
}

func (cors) Cors(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
	c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Range,Content-Disposition, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}
