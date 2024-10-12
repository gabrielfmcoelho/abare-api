package traits

import (
	"net/http"
	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
)


func ShowTraitsHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "show-traits", traits)
}