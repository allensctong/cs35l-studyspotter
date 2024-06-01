package schemas

//POST/GET FOR SIGNUP/LOGIN
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
/*
//POST FROM UPLOAD
type Upload struct {
	Username
	Image
	Caption
}

//POST FROM UPLOAD
type UploadPfp struct {
	Username
	ProfilePicture
}

//GET FROM MAIN/USER PAGE
type Post struct {
	Username string ‘json:”username”’
	Image 
	Caption string ‘json:”caption”’
	Comments # a dictionary or a similar structure with usernames as keys and their corresponding comments as values 
	LikeCount int ‘json:”like_count”’
	LikeByUser bool ‘json:”liked_by_user”’
	LikedUserList list ‘json:”liked_by_user”’
	UploadTime Time `json:"uploadtime"`
}

//GET FROM USER PAGE
type User struct {
	Username string `json:"username"`
	IsUsername bool   `json:"isusername"`
	Bio string `json:"bio"`
	Following int `json:"following"`
	Followers int `json:"followers"`
	ProfilePicture
	Posts []Post
}*/
