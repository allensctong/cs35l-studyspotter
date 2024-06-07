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
type Comment struct {
	Username string
	Text string
}

//GET FROM MAIN/USER PAGE
type Post struct {
	ID int `json:"post_id"`
	Username string `json:"username"`
	ImagePath string `json:"image_src"` 
	Caption string `json:"caption"`
	LikeCount int `json:"likes"`
	Liked bool `json:"liked"`
	LikedUserList []string `json:"liked_users"`
	Comments []Comment `json:"comments"`
}

//GET FROM USER PAGE
type UserProfile struct {
	Username string `json:"username"`
	Bio string `json:"bio"`
	FollowingCount int `json:"following_count"`
	Following []string `json:"following"`
	FollowersCount int `json:"followers_count"`
	Followers []string `json:"follower"`
	IsFriend bool `json:"isFriend"`
	ProfilePicture string `json:"pfp"`
	Posts []string `json:"posts"`
}
