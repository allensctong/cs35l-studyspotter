package schemas

import (
//	"time"
)

//POST/GET FOR SIGNUP/LOGIN
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//POST FROM UPLOAD
type Upload struct {
	Username string `json: "username"`
	//Image 
	Caption string `json: "caption"`
}

//POST FROM UPLOAD
type UploadPfp struct {
	Username string `json: "username"`
	//ProfilePicture 
}

//GET FROM MAIN/USER PAGE
type Post struct {
	Username string `json:"username"`
	ImageSrc string `json:"image_src"` 
	Caption string `json:"caption"`
	//Comments # a dictionary or a similar structure with usernames as keys and their corresponding comments as values 
	LikeCount int `json:"like_count"`
	LikeByUser bool `json:"liked_by_user"`
	LikedUserList []string `json:"liked_by_user"`
}

//GET FROM USER PAGE
type UserProfile struct {
	Username string `json:"username"`
	Bio string `json:"bio"`
	FollowingCount int `json:"following_count"`
	FollowersCount int `json:"followers_count"`
	ProfilePicture string `json:"pfp"`
}
