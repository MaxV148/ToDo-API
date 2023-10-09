package api

type BaseResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires-at"`
}

type BaseRequest struct {
	Token string `json:"token" binding:"required"`
}
