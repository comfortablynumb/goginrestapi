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

// Find Search for users.
// @Summary Search for users.
// @Description Allows you to search for users using different filters and options.
// @Produce json
// @Param username query string false "Username"
// @Param sort_by query string false "Field to sort by. Allowed fields: username"
// @Param sort_dir query string false "Direction to sort by. Allowed values: asc, desc. Default: asc"
// @Param offset query int false "Starts results from this offset. Default: 0"
// @Param limit query int false "Limits the amount of results to return. Default: 50"
// @Success 200 {object} resource.UserResource
// @Failure 400 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags users
// @Router /user [get]
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

// Create Create a new user.
// @Summary Create a new user.
// @Description Allows you to create a new user.
// @Accept json
// @Produce json
// @Param user body resource.UserCreateResource true "User data"
// @Success 201 {object} resource.UserResource
// @Failure 400 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags users
// @Router /user [post]
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

// Update Update a user.
// @Summary Update a user.
// @Description Allows you to update an existing user.
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param user body resource.UserUpdateResource true "User data"
// @Success 200 {object} resource.UserResource
// @Failure 400 {object} apperror.HttpError
// @Failure 404 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags users
// @Router /user/{username} [put]
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

// Delete Delete a user.
// @Summary Delete a user.
// @Description Allows you to delete an existing user.
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param user body resource.UserDeleteResource true "User data"
// @Success 200 {object} resource.UserResource
// @Failure 400 {object} apperror.HttpError
// @Failure 404 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags users
// @Router /user/{username} [delete]
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
