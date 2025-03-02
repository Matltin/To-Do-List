package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	mock_sqlc "to_do_list/db/mock"
	"to_do_list/db/sqlc"
	"to_do_list/token"
	"to_do_list/util"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteTodoAPI(t *testing.T) {
	util.Load("../.env")
	tokenMaker, _ := token.NewPasetoMaker(os.Getenv("PASETO_SYMMETRIC_KEY"))

	todo := randomTodo()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_sqlc.NewMockStore(ctrl)
	store.EXPECT().
		DeleteTodo(gomock.Any(), gomock.Eq(sqlc.DeleteTodoParams{
			ID:     todo.ID,
			UserID: todo.UserID,
		})).
		Times(1).
		Return(nil)

	server := NewServer(store, *tokenMaker)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/todos/%d", todo.ID) 
	request, err := http.NewRequest(http.MethodDelete, url, nil) 

	require.NoError(t, err)

	authToken, err := tokenMaker.CreateToken(uint(todo.UserID), "abcd", time.Minute)
	require.NoError(t, err)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	server.Router.ServeHTTP(recorder, request)

	fmt.Print("\n\n", authToken, "\n\n")

	require.Equal(t, http.StatusOK, recorder.Code)
}


func randomTodo() sqlc.Todo {
	return sqlc.Todo{
		ID: int64(util.RandRange(1, 100)),
		UserID: int32(util.RandRange(1, 100)),
		CreateAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		Title: "asdf",
		Description: "asdf",
		IsDone: sql.NullBool{
			Bool: false,
			Valid: true,
		},
	}
}
