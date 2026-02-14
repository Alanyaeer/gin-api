package service

import (
	"chat-system/config"
	"chat-system/internal/model/user"
	"chat-system/internal/repository"
	"chat-system/internal/repository/local/userinfo"
)

var userRepo repository.UserRepository = getUserRepo()

func GetUserInfoByUserId(userId string) user.UserInfo {
	if userInfo, err := getUserRepo().GetUserInfoByUserId(userId); err == nil {
		return *userInfo
	} else {
		return user.UserInfo{}
	}
}
func AddUserInfo(userInfo user.UserInfo) error {
	return getUserRepo().CreateUser(&userInfo)
}

func getUserRepo() repository.UserRepository {
	switch config.DbType {
	case "locaL":
		return userinfo.NewUserLocalRepo()
	case "mysql":
		return nil
	default:
		return userinfo.NewUserLocalRepo()
	}
}
