package traits

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
)


func CreateTraitsHandler(ctx *gin.Context) {
	request := CreateTraitsRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	traits := schemas.Traits{
		IsValid: *request.IsValid,
		Value:   request.Value,
		ChildID: request.ChildID,
		TagID:  request.TagID,
	}
	if err := handler.Db.Create(&traits).Error; err != nil {
		handler.Logger.Errorf("error creating traits: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating traits on database")
		return
	}
	sendSuccess(ctx, "create-traits", traits)
}