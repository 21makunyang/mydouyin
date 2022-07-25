package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mydouyin/models"
	"net/http"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}
type UserResponse struct {
	Response
	User User `json:"user"`
}

var UserLoginInfo = map[string]User{}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	newuser := NameAndPassword{}
	var count int64 = 0

	models.DB.Model(&newuser).Where("name=?", username).Count(&count)
	if count >= 1 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User already exist",
		})
	} else {
		token := username + password
		newNameAndPassword := NameAndPassword{
			Name:     username,
			Password: password,
			Token:    token,
		}
		models.DB.Create(&newNameAndPassword)

		newUser := User{
			Id:            newNameAndPassword.Id,
			Name:          username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}

		models.DB.Create(&newUser)

		//models.DB.Where("name=?", username).Find(&User{})
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newUser.Id,
		})
	}
}
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	nameAndPassword := NameAndPassword{}
	user := User{}
	if err := models.DB.Where("name=? AND password=?", username, password).Find(&nameAndPassword).Error; err == nil {
		models.DB.Where("id=?", nameAndPassword.Id).Find(&user)
		UserLoginInfo[token] = user
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   nameAndPassword.Id,
			Token:    token,
		})
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "wrong username or password"},
		})
	} else {
		fmt.Println(err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "nil"},
		})
	}
}
func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if user, exist := UserLoginInfo[token]; exist != false {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "user does not exist"},
		})
	}
}

//safty
