package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mydouyin/models"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	token := c.PostForm("token")
	NP := NameAndPassword{}
	if _, exist := UserLoginInfo[token]; exist == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}

	title := c.PostForm("title")

	filename := filepath.Base(data.Filename)
	models.DB.Where("token=?", token).Find(&NP)
	user := User{}
	models.DB.Where("id=?", NP.Id).Find(&user)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public", finalName)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}

	videoname := strconv.FormatInt(user.Id, 10) + title + ".mp4"
	models.UploadAliyunOss(videoname, saveFile)
	video := Video{
		Author:        user,
		PlayUrl:       "https://mkydouyin.oss-cn-guangzhou.aliyuncs.com/" + videoname,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		isFavorite:    false,
		Title:         title,
		PublishTime:   time.Now(),
	}
	models.DB.Model(&Video{}).Create(&video)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		//StatusMsg:  " upload successfully",
		StatusMsg: finalName + " upload successfully",
	})
}

func PublishList(c *gin.Context) {
	token := c.Query("token")
	NP := NameAndPassword{}
	if _, exist := UserLoginInfo[token]; exist == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}

	user := User{}
	models.DB.Where("token=?", token).Find(&NP)
	models.DB.Where("name=?", NP.Name).Find(&user)
	videoList := []Video{}
	models.DB.Where("author_id=?", user.Id).Find(&videoList)
	for i := range videoList {
		videoList[i].Author = user
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
