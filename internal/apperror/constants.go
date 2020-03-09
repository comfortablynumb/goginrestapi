package apperror

// Constants

const (
	InternalErrorCode    = "000001"
	InternalErrorMessage = "Internal server error."

	BindingErrorCode    = "000002"
	BindingErrorMessage = "Request binding error"

	ValidationErrorCode    = "000003"
	ValidationErrorMessage = "Request validation error"

	DbErrorCode    = "000004"
	DbErrorMessage = "An error occurred on our database"

	ModelNotFoundErrorCode    = "000005"
	ModelNotFoundErrorMessage = "The element you referenced was not found"
)
