package domain

// swagger:parameters GetOneUser
type GetUserRequest struct {
	// User ID 参数
	//
	// Required: true
	// in: path
	ID int `json:"id"`
}

type BaseResponseBodyForSwagger struct {
	// 错误码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
}

// swagger:response GetUserResponse
type GetUserResponse struct {
	// in: body
	Body struct {
		BaseResponseBodyForSwagger
		Data User `json:"data"`
	}
}
