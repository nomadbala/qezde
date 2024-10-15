package http

import (
	"github.com/gin-gonic/gin"
	"qezde/user/internal/service"
	"qezde/user/pkg/server/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Routes(r *gin.RouterGroup) {
	api := r.Group("/user")
	{
		api.GET("/", h.GetAllUsers)
		api.POST("/", h.CreateUser)
	}
}

// TODO: add swagger
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	res, err := h.userService.GetAllUsers()
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.OK(c, res)
}

func (h *UserHandler) CreateUser(c *gin.Context) {

}
