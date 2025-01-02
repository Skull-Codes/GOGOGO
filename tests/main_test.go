package tests

import (
	"encoding/json"
	"gogogo/controllers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMediaControllerOK(t *testing.T) {

	dirPath := "./uploads"
	err := os.MkdirAll(dirPath, os.ModePerm) // Crea la carpeta si no existe
	if err != nil {
		t.Fatalf("Error al crear la carpeta de prueba: %v", err)
	}

	filePath := "./uploads/testfile.txt"
	err = os.WriteFile(filePath, []byte("Este es un archivo de prueba."), 0644)
	if err != nil {
		t.Fatalf("Error al crear el archivo de prueba: %v", err)
	}

	r := gin.Default()
	r.GET("/media/*uploads", controllers.UploadServe)

	req := httptest.NewRequest(http.MethodGet, "/media/uploads/testfile.txt", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	t.Cleanup(func() {
		err := os.Remove(filePath)
		if err != nil {
			t.Fatalf("Error when trying to delete testfile.txt: %v", err)
		}
	})
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
