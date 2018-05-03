package route

import (
	"github.com/dracher/autorhvhprovison/provision"
	"github.com/gin-gonic/gin"
)

// AutoConfigMiddle is
func AutoConfigMiddle(cfg *provision.AutoConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cfg", cfg)
		c.Next()
	}
}
