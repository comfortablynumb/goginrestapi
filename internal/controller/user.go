package controller

import (
	"net/http"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/resource"
	"github.com/comfortablynumb/goginrestapi/internal/service"
	"github.com/gin-gonic/gin"
)

// Constants

const (
	UserControllerSourceName = "UserController"
)

// Structs

type UserController struct {
	userService           service.UserService
	requestContextFactory *context.RequestContextFactory
}

func (ctrl *UserController) Find(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req resource.UserFindResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserControllerSourceName, nil))

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
	var req resource.UserCreateResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserControllerSourceName, nil))

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
	var req resource.UserUpdateResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserControllerSourceName, nil))

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserControllerSourceName, nil))

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

	var req resource.UserDeleteResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserControllerSourceName, nil))

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

func NewUserController(userService service.UserService, requestContextFactory *context.RequestContextFactory) *UserController {
	return &UserController{
		userService:           userService,
		requestContextFactory: requestContextFactory,
	}
}
