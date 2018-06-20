package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index -
func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "https://github.com/zhs007/tradingview_udf_go")
	}
}
