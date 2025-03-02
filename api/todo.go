package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"to_do_list/db/sqlc"
	"to_do_list/token"
	"to_do_list/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (ser *Server) CreateTodo(ctx *gin.Context) {
	var req CreateTodoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPyloadKey).(*token.Payload)

	arg := sqlc.CreateTodoParams{
		UserID:      int32(authPayload.UserID),
		Title:       req.Title,
		Description: req.Description,
	}

	todo, err := ser.store.CreateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (ser *Server) UpdateTodo(ctx *gin.Context) {
	var req UpdateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := ctx.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if todoID < 1 {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPyloadKey).(*token.Payload)

	arg := sqlc.UpdateTodoParams{
		ID:          int64(todoID),
		UserID:      int32(authPayload.UserID),
		Title:       req.Title,
		Description: req.Description,
	}

	todo, err := ser.store.UpdateTodo(ctx, arg)
	if err != nil {
		if err == util.ErrActivityDone {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (ser *Server) DeleteTodo(ctx *gin.Context) {

	id := ctx.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPyloadKey).(*token.Payload)

	arg := sqlc.DeleteTodoParams{
		ID:     int64(todoID),
		UserID: int32(authPayload.UserID),
	}

	err = ser.store.DeleteTodo(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ser *Server) GetTodos(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))    // Default to page 1
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10")) // Default to limit 10

	if page < 1 {
		ctx.JSON(http.StatusBadRequest, errorResponse(util.ErrInvalidPageNumber))
		return
	}

	if limit < 1 {
		ctx.JSON(http.StatusBadRequest, errorResponse(util.ErrInvalidLimitNumber))
		return
	}

	authPayload := ctx.MustGet(authorizationPyloadKey).(*token.Payload)

	arg := sqlc.GetTodosByIDParams{
		UserID: int32(authPayload.UserID),
		Offset: int32(page),
		Limit:  int32(limit),
	}

	todos, err := ser.store.GetTodosByID(ctx, arg)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todos)
}
