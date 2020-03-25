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

// Find Search for user types.
// @Summary Search for user types.
// @Description Allows you to search for user types using different filters and options.
// @Produce json
// @Param name query string false "User Type Name"
// @Param sort_by query string false "Field to sort by. Allowed fields: name"
// @Param sort_dir query string false "Direction to sort by. Allowed values: asc, desc. Default: asc"
// @Param offset query int false "Starts results from this offset. Default: 0"
// @Param limit query int false "Limits the amount of results to return. Default: 50"
// @Success 200 {object} resource.UserTypeResourceList
// @Failure 400 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags user types
// @Router /user_type [get]
func (ctrl *UserTypeController) Find(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	var req resource.UserTypeFindResource

	if err := c.ShouldBind(&req); err != nil {
		c.Error(apperror.NewBindingHttpError(requestContext, err, UserTypeControllerSourceName, nil))

		return
	}

	userResourceList, err := ctrl.userTypeService.Find(requestContext, &req)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResourceList)
}

// Find a user type by its name.
// @Summary Find a user type by its name.
// @Description Allows you to search a user type by its name
// @Produce json
// @Param name path string true "User Type Name"
// @Success 200 {object} resource.UserTypeResource
// @Failure 404 {object} apperror.HttpError
// @Failure 400 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags user types
// @Router /user_type/{name} [get]
func (ctrl *UserTypeController) FindOneByName(c *gin.Context) {
	requestContext := ctrl.requestContextFactory.NewRequestContext(c)
	userResourceList, err := ctrl.userTypeService.FindOneByName(requestContext, c.Param("name"))

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, userResourceList)
}

// Create Create a new user type.
// @Summary Create a new user type.
// @Description Allows you to create a new user type.
// @Accept json
// @Produce json
// @Param user body resource.UserTypeCreateResource true "User Type data"
// @Success 201 {object} resource.UserTypeResource
// @Failure 400 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags user types
// @Router /user_type [post]
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

// Update Update a user type.
// @Summary Update a user type.
// @Description Allows you to update an existing user type.
// @Accept json
// @Produce json
// @Param name path string true "Name"
// @Param user body resource.UserTypeUpdateResource true "User Type data"
// @Success 200 {object} resource.UserTypeResource
// @Failure 400 {object} apperror.HttpError
// @Failure 404 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags user types
// @Router /user_type/{name} [put]
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

// Delete Delete a user type.
// @Summary Delete a user type.
// @Description Allows you to delete an existing user type.
// @Accept json
// @Produce json
// @Param name path string true "Name"
// @Success 200 {object} resource.UserTypeResource
// @Failure 400 {object} apperror.HttpError
// @Failure 404 {object} apperror.HttpError
// @Failure 500 {object} apperror.HttpError
// @Tags user types
// @Router /user_type/{name} [delete]
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
