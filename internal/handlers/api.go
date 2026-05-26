package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jxcf/jxcf-api/internal/config"
	"github.com/jxcf/jxcf-api/internal/models"
	"github.com/jxcf/jxcf-api/internal/service"
)

func StartServer(cfg *config.Config) error {
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	// Add CORS middleware
	r.Use(corsMiddleware())

	// Initialize analyzer service
	analyzer := service.NewAnalyzer(cfg)
	batchAnalyzer := service.NewBatchAnalyzer(cfg)

	// Health check endpoint
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Metrics endpoint
	r.GET("/api/metrics", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version":   "1.0.0",
			"service":   "jxcf-api",
			"timestamp": time.Now(),
		})
	})

	// Article analysis endpoint
	r.POST("/api/analyze", func(c *gin.Context) {
		var req models.AnalyzeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: fmt.Sprintf("Invalid request: %v", err),
				Code:  http.StatusBadRequest,
			})
			return
		}

		// Validate article
		req.Article = strings.TrimSpace(req.Article)
		if len(req.Article) < 10 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Article must be at least 10 characters",
				Code:  http.StatusBadRequest,
			})
			return
		}

		if len(req.Article) > cfg.MaxArticleLength {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: fmt.Sprintf("Article exceeds maximum length of %d", cfg.MaxArticleLength),
				Code:  http.StatusBadRequest,
			})
			return
		}

		// Default language
		if req.Language == "" {
			req.Language = "en"
		}

		resp, err := analyzer.Analyze(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: fmt.Sprintf("Analysis failed: %v", err),
				Code:  http.StatusInternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	// Batch analysis endpoint
	r.POST("/api/batch-analyze", func(c *gin.Context) {
		var req models.BatchAnalyzeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: fmt.Sprintf("Invalid request: %v", err),
				Code:  http.StatusBadRequest,
			})
			return
		}

		if len(req.Articles) == 0 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "At least one article is required",
				Code:  http.StatusBadRequest,
			})
			return
		}

		if len(req.Articles) > 100 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Batch size exceeds maximum of 100 articles",
				Code:  http.StatusBadRequest,
			})
			return
		}

		// Trim and validate articles
		for i := range req.Articles {
			req.Articles[i].Content = strings.TrimSpace(req.Articles[i].Content)
			if len(req.Articles[i].Content) < 10 || len(req.Articles[i].Content) > cfg.MaxArticleLength {
				// Validation will be done in batch analyzer for each article
			}
		}

		resp, err := batchAnalyzer.AnalyzeBatch(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: fmt.Sprintf("Batch analysis failed: %v", err),
				Code:  http.StatusInternalServerError,
			})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	// 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: "Endpoint not found",
			Code:  http.StatusNotFound,
		})
	})

	return r.Run(":" + fmt.Sprintf("%d", cfg.Port))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}