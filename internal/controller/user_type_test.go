package controller_test

import (
	"net/http"
	"testing"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/mock"
	"github.com/comfortablynumb/goginrestapi/internal/resource"
	"github.com/stretchr/testify/assert"
)

// CREATION TESTS

func TestUserTypeCreationNameIsRequired(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	req := resource.UserTypeCreateResource{}
	res := &apperror.HttpError{}

	response, err := mockApp.NewPostRequest("/user_type", mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, res.HasErrorCount(1))
	assert.True(t, res.HasErrorCountByNameAndType(1, "UserTypeCreateResource.Name", "required"))
}

func TestUserTypeCreationInvalidBinding(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	res := &apperror.HttpError{}

	response, err := mockApp.NewPostRequest("/user_type", mock.NewMockAppOptions().WithBody("{\"disabled\": \"invalid\"}").WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, res.HasErrorCount(0))
}

func TestUserTypeCreationUniqueValidation(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	req := resource.UserTypeCreateResource{
		Name:     "test-some-user-type-1",
		Disabled: false,
	}
	res := &resource.UserTypeResource{}

	response, err := mockApp.NewPostRequest("/user_type", mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Disabled, res.Disabled)

	// If we try to create it again, it should fail because of the "unique" validation

	invalidRes := &apperror.HttpError{}

	response, err = mockApp.NewPostRequest("/user_type", mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(invalidRes))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, invalidRes.HasErrorCount(1))
	assert.True(t, invalidRes.HasErrorCountByNameAndType(1, "UserTypeCreateResource.Name", "unique"))
}

func TestUserTypeCreationOk(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	CreateUserType(t, mockApp, "test-user-type-1")
}

// UPDATE TESTS

func TestUserTypeUpdateNotFound(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	req := resource.UserTypeUpdateResource{
		Name:     "some-name",
		Disabled: true,
	}
	res := &apperror.HttpError{}

	response, err := mockApp.NewPutRequest("/user_type/i-dont-exist", mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestUserTypeUpdateInvalidBinding(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	userTypeReq := CreateUserType(t, mockApp, "test-user-type-1")

	res := &apperror.HttpError{}

	response, err := mockApp.NewPutRequest("/user_type/"+userTypeReq.Name, mock.NewMockAppOptions().WithBody("{\"name\": \"asd\", \"disabled\": \"invalid\"}").WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, res.HasErrorCount(0))
}

func TestUserTypeUpdateUniqueValidation(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	userTypeReq1 := CreateUserType(t, mockApp, "test-user-type-1")
	userTypeReq2 := CreateUserType(t, mockApp, "test-user-type-2")

	req := resource.UserTypeUpdateResource{
		Name:     userTypeReq1.Name, // We should NOT be able to user this name since it already exist in another user type
		Disabled: false,
	}
	invalidRes := &apperror.HttpError{}

	response, err := mockApp.NewPutRequest("/user_type/"+userTypeReq2.Name, mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(invalidRes))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.True(t, invalidRes.HasErrorCount(1))
	assert.True(t, invalidRes.HasErrorCountByNameAndType(1, "UserTypeUpdateResource.Name", "unique"))
}

func TestUserTypeUpdateOk(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	userTypeReq := CreateUserType(t, mockApp, "test-user-type-1")

	req := resource.UserTypeUpdateResource{
		Name:     userTypeReq.Name,
		Disabled: true,
	}
	res := &resource.UserTypeResource{}

	response, err := mockApp.NewPutRequest("/user_type/"+userTypeReq.Name, mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, userTypeReq.Name, res.Name)
	assert.Equal(t, req.Disabled, res.Disabled)
}

// DELETE TESTS

func TestUserTypeDeleteAnUnexistentEntityDoesNotFailToAllowIdempotence(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	res := &resource.UserTypeResource{}

	response, err := mockApp.NewDeleteRequest("/user_type/i-dont-exist", mock.NewMockAppOptions().WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "", res.Name)
}

func TestUserTypeDeleteExistentEnttiy(t *testing.T) {
	mockApp := mock.NewMockAppWithDefaultConfig()

	defer func() {
		mockApp.App.ExecuteDbMigrationsDown()
	}()

	userTypeReq := CreateUserType(t, mockApp, "test-user-type-1")

	res := &resource.UserTypeResource{}

	response, err := mockApp.NewDeleteRequest("/user_type/"+userTypeReq.Name, mock.NewMockAppOptions().WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, userTypeReq.Name, res.Name)
}

// Helper methods

func CreateUserType(t *testing.T, mockApp *mock.MockApp, name string) *resource.UserTypeCreateResource {
	req := resource.UserTypeCreateResource{
		Name:     name,
		Disabled: false,
	}
	res := &resource.UserTypeResource{}

	response, err := mockApp.NewPostRequest("/user_type", mock.NewMockAppOptions().WithBody(req).WithExpectedResponse(res))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Disabled, res.Disabled)

	return &req
}
