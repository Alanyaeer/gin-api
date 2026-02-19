package userinfo

import (
	"chat-system/config"
	"chat-system/internal/model/entity"
	"chat-system/internal/repository"
	"chat-system/pkg/file"
	"chat-system/pkg/idgenerator"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type UserLocalRepo struct {
}

// 确保 UserLocalRepo 实现了 UserRepository 接口
var _ repository.UserRepository = (*UserLocalRepo)(nil)

func NewUserLocalRepo() repository.UserRepository {
	return &UserLocalRepo{}
}

func (r *UserLocalRepo) CreateUser(userInfo *entity.UserInfo) error {
	if userInfo == nil {
		return errors.New("userInfo input is nil")
	}
	var dbUser []entity.UserInfo

	err := file.ReadJSON(config.UserRepoJsonFilePath, &dbUser)
	if err != nil {
		return err
	}
	userInfo.UserId = idgenerator.NativeNextID()
	slog.Info(fmt.Sprintf("生成的用户ID: %d", userInfo.UserId))
	dbUser = append(dbUser, *userInfo)
	if data, err := json.Marshal(dbUser); err != nil {
		return fmt.Errorf("write file failure %v", dbUser)
	} else {
		os.WriteFile(config.UserRepoJsonFilePath, data, config.WriteFileMode)
	}
	return nil
}

func (r *UserLocalRepo) GetUserInfoByUserId(userId int64) (*entity.UserInfo, error) {
	if data, err := os.ReadFile(config.UserRepoJsonFilePath); err != nil {
		return nil, err
	} else {
		var dbUser []entity.UserInfo
		err = json.Unmarshal(data, &dbUser)
		for _, userInfo := range dbUser {
			if userInfo.UserId == userId {
				return &userInfo, nil
			}
		}
		return nil, fmt.Errorf("没有找到id为%v 的用户id", userId)
	}

}
