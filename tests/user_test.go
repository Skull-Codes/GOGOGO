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
	r.POST("/register", controllers.RegisterController)

	registerData := map[string]any{
		"username": "testuser",
		"password": "testpassword",
		"is_admin": false,
	}
	body, _ := json.Marshal(registerData)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
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
	r.POST("/login", controllers.LoginController)

	loginData := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
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
	r.POST("/change_password", controllers.ChangePasswordController)

	loginData := map[string]string{
		"username":     "testuser",
		"old_password": "testpassword",
		"new_password": "newpassword",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPost, "/change_password", bytes.NewReader(body))
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
	r.POST("/delete_user", controllers.DeleteUserController)

	loginData := map[string]string{
		"username": "testuser",
		"password": "newpassword",
	}
	body, _ := json.Marshal(loginData)
	req := httptest.NewRequest(http.MethodPost, "/delete_user", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User delete sucessfully", response["message"])
}
