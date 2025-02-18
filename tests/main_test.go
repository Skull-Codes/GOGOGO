package tests

import (
	"encoding/json"
	"gogogo/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMediaControllerOK(t *testing.T) {
	r := gin.Default()
	r.GET("/media/*uploads", controllers.UploadServe)

	req := httptest.NewRequest(http.MethodGet, "/media/uploads/testdir/testfile.txt", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMediaControllerFail(t *testing.T) {
	r := gin.Default()
	r.GET("/media/*uploads", controllers.UploadServe)

	// Crea la solicitud GET
	req := httptest.NewRequest(http.MethodGet, "/media/", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Illegal request", response["message"])
}
