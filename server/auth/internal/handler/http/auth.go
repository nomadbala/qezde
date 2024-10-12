package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "qezde/auth/docs"
	"qezde/auth/internal/service"
	"qezde/auth/pkg/server/response"
)

type AuthenticationHandler struct {
	s *service.AuthenticationService
}

func NewAuthenticationHandler(s *service.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{s}
}

func (h *AuthenticationHandler) Routes(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) {
		response.OK(c, "qezde/authentication service is healthy :)")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/sign-up", h.SignUp)
	r.POST("/sign-in", h.SignIn)
}
