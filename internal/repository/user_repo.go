package repository
import (
	"chat-system/internal/model/user"
)
type UserRepository interface {
	CreateUser(userInfo *user.UserInfo) error
	GetUserInfoByUserId(userId string) (*user.UserInfo, error)
}