package handlers

import (
	dto "moonlay-test/internal/app/dto/result"
	sublistsdto "moonlay-test/internal/app/dto/sublist"
	"moonlay-test/internal/app/models"
	"moonlay-test/internal/app/repositories"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerSubList struct {
	SubListRepository  repositories.SubListRepository
	TodoListRepository repositories.TodoListRepository
}

func HandlerSubList(SubListRepository repositories.SubListRepository, TodoListRepository repositories.TodoListRepository) *handlerSubList {
	return &handlerSubList{SubListRepository, TodoListRepository}
}

func (h *handlerSubList) GetAllSubLists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	search := c.QueryParam("search")

	sublists, totalCount, err := h.SubListRepository.GetAllSubLists(page, pageSize, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	responseData := struct {
		TotalCount int              `json:"total_count"`
		Page       int              `json:"page"`
		PageSize   int              `json:"page_size"`
		Data       []models.SubList `json:"data"`
	}{
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Data:       sublists,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "SubList data successfully obtained", Data: responseData})
}

func (h *handlerSubList) GetSubList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var sublist models.SubList
	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: sublist})
}

func (h *handlerSubList) CreateSubList(c echo.Context) error {
	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")

	todolist, _ := strconv.Atoi(c.FormValue("todo_id"))
	request := sublistsdto.CreateSubListRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
		TodoID:      todolist,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	todoid, err := h.TodoListRepository.GetTodoList(request.TodoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	sublist := models.SubList{
		Title:       request.Title,
		Description: request.Description,
		Files:       request.Files,
		TodoID:      request.TodoID,
		Todo:        convertTodoResponse(todoid),
	}

	data, err := h.SubListRepository.CreateSubList(sublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"todo_id":     data.TodoID,
		"todo":        data.Todo,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "data sudah nambah", Data: result})
}

func (h *handlerSubList) UpdateSubList(c echo.Context) error {
	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")

	todolist, _ := strconv.Atoi(c.FormValue("todo_id"))
	request := sublistsdto.UpdateSubListRequest{
		TodoID:      todolist,
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SubListRepository.GetSubList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.TodoID != 0 {
		sublist.TodoID = request.TodoID
	}

	if request.Title != "" {
		sublist.Title = request.Title
	}

	if request.Description != "" {
		sublist.Description = request.Description
	}

	if request.Files != "" {
		sublist.Files = request.Files
	}

	data, err := h.SubListRepository.UpdateSubList(sublist)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})

	}
	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"todo_id":     data.TodoID,
		"todo":        data.Todo,
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "data sudah berhasil di update", Data: result})
}

func convertTodoResponse(c models.TodoList) models.TodoListResponse {
	return models.TodoListResponse{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Files:       c.Files,
	}
}
