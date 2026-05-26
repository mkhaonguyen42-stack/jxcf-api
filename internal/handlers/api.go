package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/models"
	"github.com/jxcf/jxcf-api/internal/service"
)

func StartServer(cfg *config.Config) error {
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	// Initialize analyzer service
	analyzer := service.NewAnalyzer(cfg)

	// Health check endpoint
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Article analysis endpoint
	r.POST("/api/analyze", func(c *gin.Context) {
		var req models.AnalyzeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: err.Error(),
				Code: http.StatusBadRequest,
			})
			return
		}

		if req.Language == "" {
			req.Language = "zh"
		}

		resp, err := analyzer.Analyze(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
				Code: http.StatusInternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	return r.Run(":" + string(rune(cfg.Port)))
}
