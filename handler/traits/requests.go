package traits

import (
	"fmt"
	"github.com/gabrielfmcoelho/abare-api/schemas"
)


func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening

type CreateTraitsRequest struct {
	IsValid *bool `json:"is_valid"`
    Value   string `json:"value"`
    ChildID uint `json:"child_id"`
    TagID  []schemas.Tag `json:"tag_id"`
}

func (r *CreateTraitsRequest) Validate() error {
	if r.IsValid == nil && r.Value == "" && r.ChildID <= 0 && r.TagID == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.IsValid == nil {
		return errParamIsRequired("is_valid", "bool")
	}
	if r.Value == "" {
		return errParamIsRequired("value", "string")
	}
	if r.ChildID < 0 {
		return errParamIsRequired("child_id", "uint")
	}
	if r.TagID == nil {
		return errParamIsRequired("tag_id", "[]Tag")
	}
	return nil
}

// UpdateTraits

type UpdateTraitsRequest struct {
	IsValid *bool `json:"is_valid"`
    Value   string `json:"value"`
    ChildID uint `json:"child_id"`
    TagID  []schemas.Tag `json:"tag_id"`
}

func (r *UpdateTraitsRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.IsValid != nil || r.Value != "" || r.ChildID > 0 || r.TagID != nil{
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}