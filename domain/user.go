package domain

import (
	"context"
)

// User 实体
// swagger:model User
type User struct {
	// User ID
	ID int `json:"id"`
	// 用户名
	Name string `json:"name"`
}

type UserService interface {
	// swagger:route GET /users/{id} users GetOneUser
	//
	// 通过 ID 获取 User 信息
	//
	// 这是接口的详细描述
	//
	//     Responses:
	//       200: GetUserResponse
	Get(ctx context.Context, id int) (*User, error)
}
