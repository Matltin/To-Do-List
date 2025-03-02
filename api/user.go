package api

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"time"
	"to_do_list/db/sqlc"
	"to_do_list/util"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (ser *Server) RegisterUser(ctx *gin.Context) {
	var req RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	exists, err := ser.store.CheckUserExists(ctx, req.Email)
	if exists || err != nil {
		if exists {
			ctx.JSON(http.StatusConflict, errorResponse(errors.New("user with this email already exist")))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashedPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := sqlc.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}

	user, err := ser.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	accessTokenDurationStr := os.Getenv("ACCESS_TOKEN_DURATION")
	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, err := ser.Maker.CreateToken(uint(user.ID), user.Email, accessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := UserResponse{
		ID:       uint(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	ctx.JSON(http.StatusOK, res)
}

type LoginUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (ser *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := ser.store.GetUser(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accessTokenDurationStr := os.Getenv("ACCESS_TOKEN_DURATION")
	accessTokenDuration, err := time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, err := ser.Maker.CreateToken(uint(user.ID), user.Email, accessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := UserResponse{
		ID:       uint(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	ctx.JSON(http.StatusOK, res)
}
