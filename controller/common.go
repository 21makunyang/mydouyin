package controller

import "time"

type NameAndPassword struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token"`
}
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type Video struct {
	Id            int64     `json:"id,omitempty"`
	Author        User      `json:"author"`
	AuthorId      int64     `json:"author_id"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	isFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title,omitempty"`
	PublishTime   time.Time `json:"publish_time,omitempty"`
}
type Relation struct {
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
	User     User  `json:"user,omitempty"`
}
type Favorite struct {
	VideoId  int64  `json:"video_id"`
	UserName string `json:"user_name"`
}
