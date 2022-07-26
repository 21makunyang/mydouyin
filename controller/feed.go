package controller

import (
	"github.com/gin-gonic/gin"
	"mydouyin/models"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

//Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	NP := NameAndPassword{}
	videos := []Video{}
	if _, exist := UserLoginInfo[token]; exist == false {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	} else {
		models.DB.Order("publish_time desc").Limit(10).Find(&videos)
		for i, video := range videos {
			authorId := video.AuthorId
			author := User{}
			models.DB.Where("id=?", authorId).Find(&author)
			models.DB.Where("token=?", token).Find(&NP)
			loger := User{}
			models.DB.Where("name=?", NP.Name).Find(&loger)
			var focus int64 = 0
			models.DB.Model(&Relation{}).Where("user_id=? AND to_user_id=?", loger.Id, authorId).Count(&focus)

			if focus == 1 {
				author.IsFollow = true
			} else {
				author.IsFollow = false
			}
			videos[i].Author = author

			var like int64 = 0
			models.DB.Model(&Favorite{}).Where("video_id=? AND user_name", video.Id, loger.Name).Count(&like)
			if like == 1 {
				videos[i].isFavorite = true
			} else {
				videos[i].isFavorite = false
			}
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})

}
