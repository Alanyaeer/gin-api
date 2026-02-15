package repository
import (
	"chat-system/internal/model/entity"
)
type UserRepository interface {
	CreateUser(userInfo *entity.UserInfo) error
	GetUserInfoByUserId(userId string) (*entity.UserInfo, error)
}