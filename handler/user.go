package handler

import (
	"log"
	"net/http"
	"rumah-sakit/middleware"
	"rumah-sakit/model/dto"
	usacase "rumah-sakit/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler interface {
	GetAllUser(c *gin.Context)
	GetUser(c *gin.Context)
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func CreateUserHandler(usecase usacase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: usecase,
	}
}

type userHandler struct {
	userUseCase usacase.UserUseCase
}

func (user *userHandler) GetAllUser(c *gin.Context) {

	// Memanggil use case untuk mendapatkan semua User
	users, err := user.userUseCase.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data User dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        users,
		"Status Code": 200,
		"Message":     "Success get all User",
	})
}

func (user *userHandler) GetUser(c *gin.Context) {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid User ID:", err)
	}

	// Memanggil use case untuk mendapatkan User
	get_User, err := user.userUseCase.GetUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data User dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        get_User,
		"Status Code": 200,
		"Message":     "Success get User",
	})
}

func (user *userHandler) Login(c *gin.Context) {

	var request_user dto.User_request

	if err := c.ShouldBind(&request_user); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan User
	login_user, err := user.userUseCase.Login(request_user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claims := jwt.MapClaims{
		"user_id": login_user.Id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tkn.SignedString(middleware.PrivateKey)
	c.JSON(200, gin.H{
		"token":   token,
		"Message": "Success login user",
	})

}

func (user *userHandler) CreateUser(c *gin.Context) {

	var request_user dto.User_request

	if err := c.ShouldBind(&request_user); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan User
	create_user, err := user.userUseCase.CreateUser(request_user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data User dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        create_user,
		"Status Code": 201,
		"Message":     "Success create User",
	})
}

func (user *userHandler) UpdateUser(c *gin.Context) {

	var request_user dto.User_request

	if err := c.ShouldBind(&request_user); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	// Memanggil use case untuk mendapatkan User
	update_User, err := user.userUseCase.UpdateUser(request_user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data User dalam format JSON
	c.JSON(200, map[string]any{
		"Data":        update_User,
		"Status Code": 200,
		"Message":     "Success update User",
	})
}

func (user *userHandler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid User ID:", err)
	}

	// Memanggil use case untuk mendapatkan User
	err = user.userUseCase.DeleteUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan data User dalam format JSON
	c.JSON(200, map[string]any{
		"Status Code": 200,
		"Message":     "Success delete User",
	})
}
