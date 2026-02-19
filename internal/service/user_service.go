package service

import (
	"chat-system/config"
	"chat-system/internal/model/converter"
	"chat-system/internal/model/dto"
	"chat-system/internal/model/entity"
	"chat-system/internal/repository"
	mysql "chat-system/internal/repository/mysql/userinfo"
	local "chat-system/internal/repository/local/userinfo"
)

var userRepo repository.UserRepository = getUserRepo()

func GetUserInfoByUserId(userId int64) entity.UserInfo {
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
		return local.NewUserLocalRepo()
	case "mysql":
		return mysql.NewUserMysqlRepo()
	default:
		return local.NewUserLocalRepo()
	}
}	
