package api

import (
	"to_do_list/db"
	"to_do_list/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	DB     db.DataBase
	Router *gin.Engine
	Maker  token.Paseto
}

func NewServer(db db.DataBase, maker token.Paseto) Server {
	ser := Server{
		DB:     db,
		Router: gin.Default(),
		Maker:  maker,
	}

	setupServer(ser)

	return ser
}

func setupServer(ser Server) {
	router := ser.Router

	// /api/auth/
	router.POST("/api/auth/register", ser.RegisterUser)
	router.POST("/api/auth/login", ser.LoginUser)
	router.GET("/api/auth/me")
	router.POST("/api/auth/refresh") // for the session(refresh) token
	router.POST("/api/auth/logout")  // for the session(refresh) token

	authRouter := router.Group("/", authMiddleware(ser.Maker))

	// /api/todos
	authRouter.POST("/api/todos", ser.CreateTodo)
	authRouter.PUT("api/todos/:id", ser.UpdateTodo)
	authRouter.DELETE("api/todos/:id", ser.DeleteTodo)
	authRouter.GET("/api/todos", ser.GetTodos)
}

func (ser *Server) Start(address string) error {
	return ser.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
