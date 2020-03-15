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
	UserTypeControllerSourceName = "UserTypeController"
)

// Structs

type UserTypeController struct {
	userTypeService       service.UserTypeService
	requestContextFactory *context.RequestContextFactory
}

func (ctrl *UserTypeController) Find(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req resource.UserTypeFindResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserTypeControllerSourceName, nil))

		return
	}

	userResources, err := ctrl.userTypeService.Find(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResources)
}

func (ctrl *UserTypeController) Create(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req resource.UserTypeCreateResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserTypeControllerSourceName, nil))

		return
	}

	userResource, err := ctrl.userTypeService.Create(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusCreated, userResource)
}

func (ctrl *UserTypeController) Update(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req resource.UserTypeUpdateResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserTypeControllerSourceName, nil))

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserTypeControllerSourceName, nil))

		return
	}

	userResource, err := ctrl.userTypeService.Update(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResource)
}

func (ctrl *UserTypeController) Delete(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)

	var req resource.UserTypeDeleteResource

	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserTypeControllerSourceName, nil))

		return
	}

	userResource, err := ctrl.userTypeService.Delete(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResource)
}

// Static functions

func NewUserTypeController(userTypeService service.UserTypeService, requestContextFactory *context.RequestContextFactory) *UserTypeController {
	return &UserTypeController{
		userTypeService:       userTypeService,
		requestContextFactory: requestContextFactory,
	}
}
