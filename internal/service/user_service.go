package service

import (
	"chat-system/config"
	"chat-system/internal/model/converter"
	"chat-system/internal/model/dto"
	"chat-system/internal/model/entity"
	"chat-system/internal/repository"
	"chat-system/internal/repository/local/userinfo"
)

var userRepo repository.UserRepository = getUserRepo()

func GetUserInfoByUserId(userId string) entity.UserInfo {
	if userInfo, err := getUserRepo().GetUserInfoByUserId(userId); err == nil {
		return *userInfo
	} else {
		return entity.UserInfo{}
	}
}
func AddUserInfo(userInfo *dto.UserInfoReq) error {
	if userInfoEntity, err := converter.UserInfoDtoToEntity(userInfo); err == nil {
		return getUserRepo().CreateUser(userInfoEntity)
	} else {
		return err
	}
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
