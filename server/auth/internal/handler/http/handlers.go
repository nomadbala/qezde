package http

import "github.com/gin-gonic/gin"

// SignUp sign-up godoc
// @Summary Sign up user
// @Description Sign up user
// @Tags auth
// @Accept json
// @Produce json
// @Param registrationRequest body domain.RegistrationRequest true "Sign up request"
// @Success 200 {object} domain.UserDTO
// @Failure 400 {object} response.Object
// @Failure 500 {object} response.Object
// @Router /auth/sign-up [post]
func (h *AuthenticationHandler) SignUp(c *gin.Context) {

}

// SignIn sign-in godoc
// @Summary Sign in user
// @Description Sign in user
// @Tags auth
// @Accept json
// @Produce json
// @Param registrationRequest body domain.LoginRequest true "Sign in request"
// @Success 200 {object} domain.UserDTO
// @Failure 400 {object} response.Object
// @Failure 500 {object} response.Object
// @Router /auth/sign-in [post]
func (h *AuthenticationHandler) SignIn(c *gin.Context) {

}
