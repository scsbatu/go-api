package handlers

import (
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/scsbatu/go-api/contracts"
	"github.com/scsbatu/go-api/core/helpers"
	"github.com/scsbatu/go-api/core/middlewares"
	"github.com/scsbatu/go-api/models"
	"github.com/scsbatu/go-api/utils/ierror"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type TaskHandler struct{}

func (handler TaskHandler) Any(c echo.Context) error {
	switch c.Get("Method").(string) {
	case http.MethodPost:
		return handler.Post(c)
	case http.MethodGet:
		return handler.Get(c)
	case http.MethodPut:
		return handler.Update(c)
	}
	return RawResponse(c, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

// Post is used to create a Task
func (TaskHandler) Post(c echo.Context) error {
	var response contracts.CreateTaskResponse
	var responseCode int
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	response.Method = &method
	response.RequestID = &requestID
	c.Set("path", "Task")
	c.Set("action", strings.ToLower(method))
	req := new(contracts.CreateTaskRequest)
	if err := helpers.ExtractAndValidate(c, req); err != nil {
		responseCode = http.StatusBadRequest
		response.HTTPCode = &responseCode
		response.ErrorData = err
		e, _ := json.Marshal(err)
		log.Error(c, "[Task] ExtractAndValidate error", string(e))
		return RawResponse(c, response, responseCode)
	}
	rq, _ := json.Marshal(req)
	log.Info(c, "[Task] Got Req: ", string(rq))
	responseCode, err := createTask(c, req, &response)
	if responseCode == 0 {
		responseCode = http.StatusInternalServerError
	}
	if err != nil {
		e := &contracts.ErrorData{
			Code:        err.GetCode(),
			Description: err.Error(),
		}
		response.ErrorData = e
		log.Error(c, "[Task] Got Error:", err)
	}
	response.HTTPCode = &responseCode
	rs, _ := json.Marshal(response)
	log.Info(c, "[Task] Response: ", string(rs))
	return RawResponse(c, response, responseCode)
}

// Get is used to create a Task
func (TaskHandler) Get(c echo.Context) error {
	var response contracts.GetTaskResponse
	var responseCode int
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	response.Method = &method
	response.RequestID = &requestID
	c.Set("path", "user")
	c.Set("action", strings.ToLower(method))
	req := new(contracts.GetTaskRequest)
	TaskID := c.Param("task_id")
	if TaskID != "" {
		req.ID = &TaskID
	}
	if err := helpers.ExtractAndValidate(c, req); err != nil {
		responseCode = http.StatusBadRequest
		response.HTTPCode = &responseCode
		response.ErrorData = err
		e, _ := json.Marshal(err)
		log.Error(c, "[Task] ExtractAndValidate error", string(e))
		return RawResponse(c, response, responseCode)
	}
	rq, _ := json.Marshal(req)
	log.Info(c, "[Task] Got Req: ", string(rq))
	responseCode, err := fetchTask(c, req, &response)
	if responseCode == 0 {
		responseCode = http.StatusInternalServerError
	}
	if err != nil {
		e := &contracts.ErrorData{
			Code:        err.GetCode(),
			Description: err.Error(),
		}
		response.ErrorData = e
		log.Error(c, "[Task] Got Error:", err)
	}
	response.HTTPCode = &responseCode
	rs, _ := json.Marshal(response)
	log.Info(c, "[Task] Response: ", string(rs))
	return RawResponse(c, response, responseCode)
}

// Update is used to create a Task
func (TaskHandler) Update(c echo.Context) error {
	var response contracts.UpdateTaskResponse
	var responseCode int
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	response.Method = &method
	response.RequestID = &requestID
	c.Set("path", "Task")
	c.Set("action", strings.ToLower(method))
	req := new(contracts.UpdateTaskRequest)
	TaskID := c.Param("task_id")
	if TaskID != "" {
		req.ID = &TaskID
	}
	if err := helpers.ExtractAndValidate(c, req); err != nil {
		responseCode = http.StatusBadRequest
		response.HTTPCode = &responseCode
		response.ErrorData = err
		e, _ := json.Marshal(err)
		log.Error(c, "[Task] ExtractAndValidate error", string(e))
		return RawResponse(c, response, responseCode)
	}
	rq, _ := json.Marshal(req)
	log.Info(c, "[Task] Got Req: ", string(rq))
	responseCode, err := updateTask(c, req, &response)
	if responseCode == 0 {
		responseCode = http.StatusInternalServerError
	}
	if err != nil {
		e := &contracts.ErrorData{
			Code:        err.GetCode(),
			Description: err.Error(),
		}
		response.ErrorData = e
		log.Error(c, "[Task] Got Error:", err)
	}
	response.HTTPCode = &responseCode
	rs, _ := json.Marshal(response)
	log.Info(c, "[Task] Response: ", string(rs))
	return RawResponse(c, response, responseCode)
}

func createTask(
	c echo.Context,
	req *contracts.CreateTaskRequest,
	resp *contracts.CreateTaskResponse,
) (
	int,
	ierror.IError,
) {
	// Check if creator, provider and category exist
	if _, err := models.GetUserByID(*req.CreatorID); err != nil && err == gorm.ErrRecordNotFound {
		return http.StatusBadRequest, middlewares.ErrStatusBadRequest("Creator doesn't exist")
	} else if err != nil {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	}
	if _, err := models.GetTaskCategoryByID(*req.CategoryID); err != nil && err == gorm.ErrRecordNotFound {
		return http.StatusBadRequest, middlewares.ErrStatusBadRequest("Category doesn't exist")
	} else if err != nil {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	}

	j, err := models.CreateTask(req.Title, req.TaskKey, req.Details, req.CreatorID, req.CategoryID, req.Status)
	if err != nil {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	}
	id := hex.EncodeToString(*j.ID)
	id1, _ := uuid.Parse(id)
	idInStr := id1.String()
	d := contracts.CreateTaskData{
		ID: &idInStr,
	}
	resp.Data = &d
	return http.StatusOK, nil
}
func fetchTask(
	c echo.Context,
	req *contracts.GetTaskRequest,
	resp *contracts.GetTaskResponse,
) (
	int,
	ierror.IError,
) {
	j, err := models.GetTaskByID(*req.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	} else if err != nil {
		return http.StatusNotFound, middlewares.ErrStatusNotFound("Record not found")
	}
	var id, creatorID, categoryID string
	id1 := hex.EncodeToString(*j.ID)
	id = uuid.Must(uuid.Parse(id1)).String()
	id1 = hex.EncodeToString(*j.CreatorID)
	creatorID = uuid.Must(uuid.Parse(id1)).String()
	id1 = hex.EncodeToString(*j.CategoryID)
	categoryID = uuid.Must(uuid.Parse(id1)).String()
	d := contracts.GetTaskData{
		ID:         &id,
		Title:      j.Title,
		TaskKey:    j.TaskKey,
		Details:    j.Details,
		Status:     j.Status,
		CreatorID:  &creatorID,
		CategoryID: &categoryID,
	}
	resp.Data = &d
	return http.StatusOK, nil
}
func updateTask(
	c echo.Context,
	req *contracts.UpdateTaskRequest,
	resp *contracts.UpdateTaskResponse,
) (
	int,
	ierror.IError,
) {
	j, err := models.GetTaskByID(*req.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	} else if err != nil {
		return http.StatusNotFound, middlewares.ErrStatusNotFound("Record not found")
	}
	if req.Title != nil {
		j.Title = req.Title
	}
	if req.TaskKey != nil {
		j.TaskKey = req.TaskKey
	}
	if req.Details != nil {
		j.Details = req.Details
	}
	if err = models.UpdateTask(j); err != nil {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	}
	success := "true"
	d := contracts.UpdateTaskData{
		Success: &success,
	}
	resp.Data = &d
	return http.StatusOK, nil
}
