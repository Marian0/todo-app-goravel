package controllers

import (
	"goravel/app/helpers"
	"goravel/app/http/dtos"
	"goravel/app/http/requests/todo_requests"
	"goravel/app/models"
	"log"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type TodosController struct {
	//Dependent services
}

func NewTodosController() *TodosController {
	return &TodosController{
		//Inject services
	}
}

// GET /todos
func (c *TodosController) Index(ctx http.Context) {
	// user from context
	user := ctx.Value("user").(models.User)

	var todos []models.Todo
	err := facades.Orm().Query().Where("user_id = ?", user.ID).With("User").Find(&todos)

	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondSuccess(ctx, http.StatusAccepted, dtos.TodoArrayToDTO(todos))
}

// GET /todos/{id}
func (c *TodosController) Show(ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"feature": "coming soon...",
	})
}

// POST /todos
func (c *TodosController) Store(ctx http.Context) {
	// user from context
	user := ctx.Value("user").(models.User)

	// validate input
	validator, err := ctx.Request().Validate(map[string]string{
		"title": "required|max_len:200",
	})
	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if validator.Fails() {
		helpers.RespondError(ctx, http.StatusUnprocessableEntity, validator.Errors().All())
		return
	}

	var todo models.Todo
	if err := validator.Bind(&todo); err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	todo.UserID = user.ID

	if err := facades.Orm().Query().Create(&todo); err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = facades.Orm().Query().Load(&todo, "User")

	if err != nil {
		log.Println(err)
	}

	helpers.RespondSuccess(ctx, http.StatusCreated, dtos.TodoToDTO(todo))
}

// PUT /todos/{id}
func (c *TodosController) Update(ctx http.Context) {
	// Validate querystring
	todoID, err := helpers.ValidateUUID(ctx.Request().Input("id"))

	if err != nil {
		helpers.RespondError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get entity
	var todo models.Todo
	err = facades.Orm().Query().Where("id = ?", todoID).FindOrFail(&todo)
	if err != nil {
		helpers.RespondError(ctx, http.StatusNotFound, "Todo doesnt exist")
		return
	}

	// Policy check
	if !facades.Gate().WithContext(ctx).Allows("update-todo", map[string]any{
		"todo": todo,
	}) {
		helpers.RespondError(ctx, http.StatusForbidden, "Todo forbidden access")
		return
	}

	// Request validation
	var updateTodoRequest todo_requests.UpdatedTodoRequest
	errors, err := ctx.Request().ValidateRequest(&updateTodoRequest)

	// Server errors
	if err != nil {
		log.Println(err.Error())
		helpers.RespondError(ctx, http.StatusInternalServerError, "validation error")
		return
	}

	// Request errors
	if errors != nil {
		helpers.RespondError(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}

	//update model
	todo.Title = updateTodoRequest.Title
	err = facades.Orm().Query().Save(&todo)

	if err != nil {
		log.Println(err.Error())
		helpers.RespondError(ctx, http.StatusInternalServerError, "saving error")
		return
	}

	helpers.RespondSuccess(ctx, http.StatusOK, dtos.TodoToDTO(todo))
}

// DELETE /todos/{id}
func (c *TodosController) Destroy(ctx http.Context) {
	// Validate querystring
	todoID, err := helpers.ValidateUUID(ctx.Request().Input("id"))

	if err != nil {
		helpers.RespondError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Get entity
	var todo models.Todo
	err = facades.Orm().Query().Where("id = ?", todoID).FindOrFail(&todo)
	if err != nil {
		helpers.RespondError(ctx, http.StatusNotFound, "Todo doesnt exist")
		return
	}

	// Policy check
	if !facades.Gate().WithContext(ctx).Allows("destroy-todo", map[string]any{
		"todo": todo,
	}) {
		helpers.RespondError(ctx, http.StatusForbidden, "Todo forbidden access")
		return
	}

	_, err = facades.Orm().Query().Delete(&todo)

	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, "delete error")
		return
	}

	helpers.RespondSuccess(ctx, http.StatusOK, nil)
}
