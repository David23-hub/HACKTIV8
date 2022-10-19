package request


type CreatePhotoReq struct {
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId int `json:"user_id"`
}