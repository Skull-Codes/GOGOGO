package controllers

import (
	"context"
	"encoding/json"
	"gogogo/db"
	"gogogo/forms"
	"gogogo/models"
	"gogogo/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func getUserCollection() *mongo.Collection {
	return db.GetCollection("users")
}

// RegisterController gestiona el registro de un nuevo usuario.
// @Summary Registro de usuario
// @Description Permite registrar un nuevo usuario en el sistema.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param data body forms.RegisterUser true "Cuerpo de la solicitud"
// @Success 201 {object} map[string]string{message=string} "Registro exitoso" example({"message": "username register successfully"})
// @Failure 400 {object} map[string]string{message=string} "Datos inválidos" example({"message": "Invalid request data"})
// @Failure 403 {object} map[string]string{message=string} "Usuario existente" example({"message": "The username already exists"})
// @Failure 404 {object} map[string]string{message=string} "Error al registrar" example({"message": "Error when trying to register user"})
// @Router /api/register [post]
func RegisterController(c *gin.Context) {
	var user forms.RegisterUser
	userCollection := getUserCollection()

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, _ := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if count > 0 {
		c.JSON(403, gin.H{
			"message": "The username already exists",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Error when trying to register user",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "username register successfully",
	})
}

// LoginController gestiona el inicio de sesión de un nuevo usuario.
// @Summary Login de usuario
// @Description Inicia sesión en la plataforma y te retorna un token.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param data body forms.LoginUser true "Cuerpo de la solicitud"
// @Success 200 {object} map[string]string{message=string} "Genera token" example({"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzUyOTI1NTksInVzZXJuYW1lIjoidGVzdHVzZXIifQ.41PX1i1qWhlidH28a3DF9-8722cnhCXv3b2fNIDvxqA"})
// @Failure 400 {object} map[string]string{message=string} "Datos inválidos" example({"message": "Invalid request data"})
// @Failure 401 {object} map[string]string{message=string} "Usuario o contraseña incorrecto" example({"message": "Incorrect username or password"})
// @Failure 500 {object} map[string]string{message=string} "Error al generar token" example({"message": "Error when trying generate token"})
// @Router /api/login [post]
func LoginController(c *gin.Context) {
	var user forms.LoginUser
	userCollection := getUserCollection()

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storedUser := models.User{}
	err = userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&storedUser)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	token, err := services.GenerateJWT(storedUser.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error when trying generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// ChangePasswordController gestiona el cambio de contraseña de usuario.
// @Summary Cambio de contraseña
// @Description Realiza un cambio de contraseña en tu usuario.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param data body forms.ChangePasswordUser true "Cuerpo de la solicitud"
// @Success 200 {object} map[string]string{message=string} "Cambio de contraseña" example({"message": "Change password successfully"})
// @Failure 400 {object} map[string]string{message=string} "Datos inválidos" example({"message": "Invalid request data"})
// @Failure 401 {object} map[string]string{message=string} "Usuario o contraseña incorrecto" example({"message": "Incorrect username or password"})
// @Router /api/change_password [patch]
func ChangePasswordController(c *gin.Context) {

	var user forms.ChangePasswordUser

	userCollection := getUserCollection()

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storedUser := models.User{}
	err = userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&storedUser)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.OldPassword))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.NewPassword), bcrypt.DefaultCost)
	user.NewPassword = string(hashedPassword)

	update := bson.M{
		"$set": bson.M{
			"password": user.NewPassword,
		},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"username": storedUser.Username}, update)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Change password successfully",
	})
}

// DeleteUserController elimina un usuario.
// @Summary Eliminar usuario
// @Description Elimina un usuario a partir de usuario y contraseña.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param data body forms.LoginUser true "Cuerpo de la solicitud"
// @Success 200 {object} map[string]string{message=string} "Usuario borrado con éxito" example({"message": "User delete sucessfully"})
// @Failure 400 {object} map[string]string{message=string} "Datos inválidos" example({"message": "Invalid request data"})
// @Failure 401 {object} map[string]string{message=string} "Usuario o contraseña incorrecto" example({"message": "Incorrect username or password"})
// @Failure 403 {object} map[string]string{message=string} "Error al intentar borrar usuario" example({"message": "Error when trying to delete user"})
// @Failure 404 {object} map[string]string{message=string} "Usuario a eliminar no encontrado" example({"message": "User not found to delete"})
// @Router /api/delete_user [post]
func DeleteUserController(c *gin.Context) {

	var user forms.LoginUser
	userCollection := getUserCollection()

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storedUser := models.User{}
	err = userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&storedUser)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username or password",
		})
		return
	}

	deleteResult, err := userCollection.DeleteOne(ctx, bson.M{"username": storedUser.Username})
	if err != nil {
		c.JSON(403, gin.H{
			"message": "Error when trying to delete user",
		})
		return
	}

	if deleteResult.DeletedCount == 0 {
		c.JSON(404, gin.H{
			"message": "User not found to delete",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User delete sucessfully",
	})
}
