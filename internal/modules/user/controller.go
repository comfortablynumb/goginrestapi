package user

import (
	"net/http"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/gin-gonic/gin"
)

// Constants

const (
	ControllerSourceName = "UserController"
)

// Structs

type UserController struct {
	userService UserService
}

func (ctrl *UserController) Find(c *gin.Context) {
	var req UserFindResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(err, ControllerSourceName))

		return
	}

	userResources, err := ctrl.userService.Find(&req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResources)
}

func (ctrl *UserController) Create(c *gin.Context) {
	var req UserCreateResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(err, ControllerSourceName))

		return
	}

	userResource, err := ctrl.userService.Create(&req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusCreated, userResource)
}

func (ctrl *UserController) Update(c *gin.Context) {
	var req UserUpdateResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(err, ControllerSourceName))

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(err, ControllerSourceName))

		return
	}

	userResource, err := ctrl.userService.Update(&req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResource)
}

func (ctrl *UserController) Delete(c *gin.Context) {
	var req UserDeleteResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(err, ControllerSourceName))

		return
	}

	userResource, err := ctrl.userService.Delete(&req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResource)
}

// Static functions

func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
