package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"gogogo/controllers"
	"gogogo/db"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func deleteUser() {
	users := db.GetCollection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := users.DeleteOne(ctx, bson.M{"username": "testuser"})
	if err != nil {
		panic("Error when trying to delete test user")
	}
}

func TestRegisterController(t *testing.T) {
	db.ConnectDB()
	deleteUser()

	r := gin.Default()
	r.POST("/api/register", controllers.RegisterController)

	registerData := map[string]any{
		"username": "testuser",
		"password": "testpassword",
		"is_admin": false,
	}
	body, _ := json.Marshal(registerData)
	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "username register successfully", response["message"])
}

func TestLoginController(t *testing.T) {

	r := gin.Default()
	r.POST("/api/login", controllers.LoginController)

	loginData := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"], "The token value should not be empty")
}

func TestChangePasswordController(t *testing.T) {

	r := gin.Default()
	r.PATCH("/api/change_password", controllers.ChangePasswordController)

	loginData := map[string]string{
		"username":     "testuser",
		"old_password": "testpassword",
		"new_password": "newpassword",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPatch, "/api/change_password", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Change password successfully", response["message"])
}

func TestDeleteUserController(t *testing.T) {

	r := gin.Default()
	r.DELETE("/api/delete_user", controllers.DeleteUserController)

	loginData := map[string]string{
		"username": "testuser",
		"password": "newpassword",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodDelete, "/api/delete_user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User delete sucessfully", response["message"])
}
