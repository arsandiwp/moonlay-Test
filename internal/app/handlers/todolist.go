package handlers

import (
	"fmt"
	dto "moonlay-test/internal/app/dto/result"
	todolistsdto "moonlay-test/internal/app/dto/todolist"
	"moonlay-test/internal/app/models"
	"moonlay-test/internal/app/repositories"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerTodoList struct {
	TodoListRepository repositories.TodoListRepository
}

func HandlerTodoList(TodoListRepository repositories.TodoListRepository) *handlerTodoList {
	return &handlerTodoList{TodoListRepository}
}

func (h *handlerTodoList) FindTodoList(c echo.Context) error {
	todos, err := h.TodoListRepository.FindTodoList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if len(todos) > 0 {
		return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Data for all todos was successfully obtained", Data: todos})
	} else {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "No record found"})
	}
}

func (h *handlerTodoList) GetAllTodoLists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	search := c.QueryParam("search")

	todolists, totalCount, err := h.TodoListRepository.GetAllTodoLists(page, pageSize, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	responseData := struct {
		TotalCount int               `json:"total_count"`
		Page       int               `json:"page"`
		PageSize   int               `json:"page_size"`
		Data       []models.TodoList `json:"data"`
	}{
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		Data:       todolists,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Todo data successfully obtained", Data: responseData})
}

func (h *handlerTodoList) GetTodoList(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var todolist models.TodoList
	todolist, err := h.TodoListRepository.GetTodoList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: todolist})
}

func (h *handlerTodoList) CreateTodoList(c echo.Context) error {
	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")

	request := todolistsdto.CreateTodoListRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	todolist := models.TodoList{
		Title:       request.Title,
		Description: request.Description,
		Files:       request.Files,
	}

	data, err := h.TodoListRepository.CreateTodoList(todolist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	fmt.Println(data.Files, "dataFiles")
	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"sublist":     data.SubList,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: result})
}

func (h *handlerTodoList) UpdateTodoList(c echo.Context) error {
	dataFile := c.Get("dataFiles").([]string)
	stringDataFiles := strings.Join(dataFile, ",")

	request := todolistsdto.UpdateTodoListRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
		Files:       stringDataFiles,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	todolist, err := h.TodoListRepository.GetTodoList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		todolist.Title = request.Title
	}

	if request.Description != "" {
		todolist.Description = request.Description
	}

	if request.Files != "" {
		todolist.Files = request.Files
	}

	data, err := h.TodoListRepository.UpdateTodoList(todolist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	fmt.Println(data.Files)
	result := map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"files":       strings.Split(data.Files, ","),
		"sublist":     data.SubList,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: result})
}

func (h *handlerTodoList) DeleteTodoList(c echo.Context) error {
	id, _ := strconv.Atoi("id")
	todolist, err := h.TodoListRepository.GetTodoList(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TodoListRepository.DeleteTodoList(todolist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
