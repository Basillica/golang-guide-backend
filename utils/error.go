package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "optional":
		return ""
	case "email":
		return "this field must be a valid email addresss"
	case "oneof":
		return "thie field must be one of the following: ADMIN, USER, ETC"
	case "alpha":
		return "this field contain only english alphabets"
	case "lte":
		return "this field should be less than " + fe.Param()
	case "gte":
		return "this field should be greater than " + fe.Param()
	case "min":
		return "this field should not be less than " + fe.Param()
	case "max":
		return "this field should not be greater than " + fe.Param()
	}
	return "this field is missing and must be provided"
}

func FormatError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{
				Field:   fe.Field(),
				Message: GetErrorMsg(fe),
			}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": out,
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
