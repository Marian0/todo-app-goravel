package controllers

import (
	"goravel/app/http/dtos"
	"goravel/app/models"

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
	err := facades.Orm.Query().Where("user_id = ?", user.ID).Find(&todos)

	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}
	ctx.Response().Success().Json(dtos.TodoToTodoDTO(todos))
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
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
		return
	}
	if validator.Fails() {
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": validator.Errors().All(),
		})
		return
	}

	var todo models.Todo
	if err := validator.Bind(&todo); err != nil {
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
		return
	}
	todo.UserID = user.ID

	if err := facades.Orm.Query().Create(&todo); err != nil {
		//@todo: implement proper error handler to hide db errors
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

	ctx.Response().Success().Json(http.Json{
		"ID": todo.ID,
	})
}

// PUT /todos/{id}
func (c *TodosController) Update(ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"feature": "coming soon...",
	})
}

// DELETE /todos/{id}
func (c *TodosController) Destroy(ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"feature": "coming soon...",
	})
}
