package handlers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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

type UserHandler struct{}

func (handler UserHandler) Any(c echo.Context) error {
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

// Post is used to create a user
func (UserHandler) Post(c echo.Context) error {
	var response contracts.CreateUserResponse
	var responseCode int
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	response.Method = &method
	response.RequestID = &requestID
	c.Set("path", "user")
	c.Set("action", strings.ToLower(method))
	req := new(contracts.CreateUserRequest)
	if err := helpers.ExtractAndValidate(c, req); err != nil {
		responseCode = http.StatusBadRequest
		response.HTTPCode = &responseCode
		response.ErrorData = err
		e, _ := json.Marshal(err)
		log.Error(c, "[user] ExtractAndValidate error", string(e))
		return RawResponse(c, response, responseCode)
	}
	rq, _ := json.Marshal(req)
	log.Info(c, "[user] Got Req: ", string(rq))
	responseCode, err := createUser(c, req, &response)
	if responseCode == 0 {
		responseCode = http.StatusInternalServerError
	}
	if err != nil {
		e := &contracts.ErrorData{
			Code:        err.GetCode(),
			Description: err.Error(),
		}
		response.ErrorData = e
		log.Error(c, "[user] Got Error:", err)
	}
	response.HTTPCode = &responseCode
	rs, _ := json.Marshal(response)
	log.Info(c, "[user] Response: ", string(rs))
	return RawResponse(c, response, responseCode)
}

// Get is used to create a user
func (UserHandler) Get(c echo.Context) error {
	var response contracts.GetUserResponse
	var responseCode int
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	response.Method = &method
	response.RequestID = &requestID
	c.Set("path", "user")
	c.Set("action", strings.ToLower(method))
	req := new(contracts.GetUserRequest)
	userID := c.Param("user_id")
	fmt.Println("params", c.ParamNames(), c.ParamValues())
	if userID != "" {
		req.ID = &userID
	}
	if err := helpers.ExtractAndValidate(c, req); err != nil {
		responseCode = http.StatusBadRequest
		response.HTTPCode = &responseCode
		response.ErrorData = err
		e, _ := json.Marshal(err)
		log.Error(c, "[user] ExtractAndValidate error", string(e))
		return RawResponse(c, response, responseCode)
	}
	rq, _ := json.Marshal(req)
	log.Info(c, "[user] Got Req: ", string(rq))
	responseCode, err := fetchUser(c, req, &response)
	if responseCode == 0 {
		responseCode = http.StatusInternalServerError
	}
	if err != nil {
		e := &contracts.ErrorData{
			Code:        err.GetCode(),
			Description: err.Error(),
		}
		response.ErrorData = e
		log.Error(c, "[user] Got Error:", err)
	}
	response.HTTPCode = &responseCode
	rs, _ := json.Marshal(response)
	log.Info(c, "[user] Response: ", string(rs))
	return RawResponse(c, response, responseCode)
}

// Update is used to create a user
func (UserHandler) Update(c echo.Context) error {
	var response contracts.UpdateUserResponse
	var responseCode int
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	response.Method = &method
	response.RequestID = &requestID
	c.Set("path", "user")
	c.Set("action", strings.ToLower(method))
	req := new(contracts.UpdateUserRequest)
	userID := c.Param("user_id")
	if userID != "" {
		req.ID = &userID
	}
	if err := helpers.ExtractAndValidate(c, req); err != nil {
		responseCode = http.StatusBadRequest
		response.HTTPCode = &responseCode
		response.ErrorData = err
		e, _ := json.Marshal(err)
		log.Error(c, "[user] ExtractAndValidate error", string(e))
		return RawResponse(c, response, responseCode)
	}
	rq, _ := json.Marshal(req)
	log.Info(c, "[user] Got Req: ", string(rq))
	responseCode, err := updateUser(c, req, &response)
	if responseCode == 0 {
		responseCode = http.StatusInternalServerError
	}
	if err != nil {
		e := &contracts.ErrorData{
			Code:        err.GetCode(),
			Description: err.Error(),
		}
		response.ErrorData = e
		log.Error(c, "[user] Got Error:", err)
	}
	response.HTTPCode = &responseCode
	rs, _ := json.Marshal(response)
	log.Info(c, "[user] Response: ", string(rs))
	return RawResponse(c, response, responseCode)
}

func createUser(
	c echo.Context,
	req *contracts.CreateUserRequest,
	resp *contracts.CreateUserResponse,
) (
	int,
	ierror.IError,
) {

	u, err := models.CreateUser(req.FirstName, req.LastName, req.DocumentNotes)
	if err != nil {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	}
	id := hex.EncodeToString(*u.ID)
	idInStr := uuid.Must(uuid.Parse(id)).String()
	d := contracts.CreateUserData{
		ID: &idInStr,
	}
	resp.Data = &d
	return http.StatusOK, nil

}
func fetchUser(
	c echo.Context,
	req *contracts.GetUserRequest,
	resp *contracts.GetUserResponse,
) (
	int,
	ierror.IError,
) {
	u, err := models.GetUserByID(*req.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	} else if err != nil {
		return http.StatusNotFound, middlewares.ErrStatusNotFound("record not found")
	}
	id := hex.EncodeToString(*u.ID)
	idInStr := uuid.Must(uuid.Parse(id)).String()

	d := contracts.GetUserData{
		ID:            &idInStr,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		DocumentNotes: u.DocumentNotes,
	}
	resp.Data = &d
	return http.StatusOK, nil

}
func updateUser(
	c echo.Context,
	req *contracts.UpdateUserRequest,
	resp *contracts.UpdateUserResponse,
) (
	int,
	ierror.IError,
) {
	u, err := models.GetUserByID(*req.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	} else if err != nil {
		return http.StatusNotFound, middlewares.ErrStatusNotFound("record not found")
	}
	if req.FirstName != nil {
		u.FirstName = req.FirstName
	}
	if req.LastName != nil {
		u.LastName = req.LastName
	}

	if req.DocumentNotes != nil {
		u.DocumentNotes = req.DocumentNotes
	}
	if err = models.UpdateUser(u); err != nil {
		return http.StatusInternalServerError, middlewares.ErrStatusInternalServerError("Database error", err)
	}
	success := "true"
	d := contracts.UpdateUserData{
		Success: &success,
	}
	resp.Data = &d
	return http.StatusOK, nil

}
