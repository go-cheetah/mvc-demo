package health

import (
	"net/http"

	"demo/internal/config"
	"demo/internal/pkg/database"

	"github.com/gin-gonic/gin"
)

func HealthHandler(conf *config.Config, db database.DB) func(c *gin.Context) {

	return func(c *gin.Context) {
		service := Health{
			Config: conf,
			DB:     db,
		}
		res := service.Status()
		c.JSON(http.StatusOK, res)
	}
}
