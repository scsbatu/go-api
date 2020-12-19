package contracts

type CreateUserRequest struct {
	BaseRequest
	FirstName     *string `json:"first_name" validate:"required"`
	LastName      *string `json:"last_name"`
	DocumentNotes *string `json:"document_notes"`
}

type CreateUserResponse struct {
	BaseResponse
	Data *CreateUserData `json:"data"`
}

type CreateUserData struct {
	ID *string `json:"id"`
}

type GetUserRequest struct {
	BaseRequest
	ID *string `path:"id" validate:"required"`
}

type GetUserResponse struct {
	BaseResponse
	Data *GetUserData `json:"data"`
}

type GetUserData struct {
	ID            *string `json:"id"`
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	DocumentNotes *string `json:"document_notes"`
}

type UpdateUserRequest struct {
	BaseRequest
	ID            *string `path:"id" validate:"required"`
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	DocumentNotes *string `json:"document_notes"`
}

type UpdateUserResponse struct {
	BaseResponse
	Data *UpdateUserData `json:"data"`
}

type UpdateUserData struct {
	Success *string `json:"success"`
}
