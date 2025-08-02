package handlers

import (
	"net/http"

	"github.com/Takumikimura1899/openapi-schema-drive-go-next/api/generated"
	"github.com/gin-gonic/gin"
)

// APIHandler implements the generated ServerInterface
type APIHandler struct{}

// NewAPIHandler creates a new API handler
func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

// GetUsers implements the generated ServerInterface
func (h *APIHandler) GetUsers(c *gin.Context) {
	// サンプルデータを返す
	users := []generated.User{
		{
			Id:   stringPtr("1"),
			Name: stringPtr("John Doe"),
		},
		{
			Id:   stringPtr("2"),
			Name: stringPtr("Jane Smith"),
		},
	}

	c.JSON(http.StatusOK, users)
}

// stringPtrはポインタヘルパー関数
func stringPtr(s string) *string {
	return &s
}
