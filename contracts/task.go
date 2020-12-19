package contracts

type CreateTaskRequest struct {
	BaseRequest
	Title      *string `json:"title" validate:"required"`
	TaskKey    *string `json:"Task_key"`
	Details    *string `json:"details"`
	Status     *int    `json:"status"`
	CreatorID  *string `json:"creator_id" validate:"required"`
	CategoryID *string `json:"category_id" validate:"required"`
}

type CreateTaskResponse struct {
	BaseResponse
	Data *CreateTaskData `json:"data"`
}

type CreateTaskData struct {
	ID *string `json:"id"`
}

type GetTaskRequest struct {
	BaseRequest
	ID *string `path:"id" validate:"required"`
}

type GetTaskResponse struct {
	BaseResponse
	Data *GetTaskData `json:"data"`
}

type GetTaskData struct {
	ID         *string `json:"id"`
	Title      *string `json:"title"`
	TaskKey    *string `json:"Task_key"`
	Details    *string `json:"details"`
	Status     *int    `json:"status"`
	CreatorID  *string `json:"creator_id"`
	ProviderID *string `json:"provider_id"`
	CategoryID *string `json:"category_id"`
}

type UpdateTaskRequest struct {
	BaseRequest
	ID      *string `path:"id" validate:"required"`
	Title   *string `json:"title"`
	TaskKey *string `json:"Task_key"`
	Details *string `json:"details"`
	Status  *int    `json:"status"`
}

type UpdateTaskResponse struct {
	BaseResponse
	Data *UpdateTaskData `json:"data"`
}

type UpdateTaskData struct {
	Success *string `json:"success"`
}
