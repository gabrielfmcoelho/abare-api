package traits

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
)




func UpdaterTraitsHandler(ctx *gin.Context) {
	request := UpdateTraitsRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	traits := schemas.Traits{}

	if err := handler.Db.First(&traits, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "traits not found")
		return
	}
	// Update traits
	if request.IsValid != nil {
		traits.IsValid = request.IsValid
	}

	if request.Value != "" {
		traits.Value = request.Value
	}

	if request.ChildID > 0 {
		traits.ChildID = request.ChildID
	}

	if request.TagID != nil {
		traits.TagID = request.TagID
	}
	// Save traits
	if err := handler.Db.Save(&traits).Error; err != nil {
		handler.Logger.Errorf("error updating traits: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating traits on database")
		return
	}
	sendSuccess(ctx, "update-traits", traits)
}