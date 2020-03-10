package user

import (
	"net/http"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/gin-gonic/gin"
)

// Constants

const (
	ControllerSourceName = "UserController"
)

// Structs

type UserController struct {
	userService           UserService
	requestContextFactory *context.RequestContextFactory
}

func (ctrl *UserController) Find(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req UserFindResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, ControllerSourceName))

		return
	}

	userResources, err := ctrl.userService.Find(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResources)
}

func (ctrl *UserController) Create(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req UserCreateResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, ControllerSourceName))

		return
	}

	userResource, err := ctrl.userService.Create(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusCreated, userResource)
}

func (ctrl *UserController) Update(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req UserUpdateResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, ControllerSourceName))

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, ControllerSourceName))

		return
	}

	userResource, err := ctrl.userService.Update(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResource)
}

func (ctrl *UserController) Delete(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)

	var req UserDeleteResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, ControllerSourceName))

		return
	}

	userResource, err := ctrl.userService.Delete(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResource)
}

// Static functions

func NewUserController(userService UserService, requestContextFactory *context.RequestContextFactory) *UserController {
	return &UserController{
		userService:           userService,
		requestContextFactory: requestContextFactory,
	}
}
