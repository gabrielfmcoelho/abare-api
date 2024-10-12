package user

import (
	"github.com/gabrielfmcoelho/abare-api/schemas"
)

type HandlerUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

type HandlerUsersResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.UserResponse `json:"data"`
}
