package userinfo

import (
	"chat-system/config"
	"chat-system/internal/model/user"
	"chat-system/internal/repository"
	"chat-system/pkg/file"
	"chat-system/pkg/idgenerator"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type UserLocalRepo struct {
}

// 确保 UserLocalRepo 实现了 UserRepository 接口
var _ repository.UserRepository = (*UserLocalRepo)(nil)

func NewUserLocalRepo() repository.UserRepository {
	return &UserLocalRepo{}
}

func (r *UserLocalRepo) CreateUser(userInfo *user.UserInfo) error {
	if userInfo == nil {
		return errors.New("userInfo input is nil")
	}
	var dbUser []user.UserInfo

	err := file.ReadJSON(config.UserRepoJsonFilePath, &dbUser)
	if err != nil {
		return err
	}
	userInfo.UserId = idgenerator.NextID()
	log.Printf("生成的用户ID: %s", userInfo.UserId)
	dbUser = append(dbUser, *userInfo)
	if data, err := json.Marshal(dbUser); err != nil {
		return fmt.Errorf("write file failure %v", dbUser)
	} else {
		os.WriteFile(config.UserRepoJsonFilePath, data, config.WriteFileMode)
	}
	return nil
}

func (r *UserLocalRepo) GetUserInfoByUserId(userId string) (*user.UserInfo, error) {
	if data, err := os.ReadFile(config.UserRepoJsonFilePath); err != nil {
		return nil, err
	} else {
		var dbUser []user.UserInfo
		err = json.Unmarshal(data, &dbUser)
		for _, userInfo := range dbUser {
			if userInfo.UserId == userId {
				return &userInfo, nil
			}
		}
		return nil, fmt.Errorf("没有找到id为%v 的用户id", userId)
	}

}
