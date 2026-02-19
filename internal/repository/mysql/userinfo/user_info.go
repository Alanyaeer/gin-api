package userinfo

import (
	"chat-system/config"
	"chat-system/internal/model/entity"
	"chat-system/internal/repository"
	"chat-system/pkg/idgenerator"
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log/slog"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(config.MysqlDsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

type UserMysqlRepo struct {
	db *gorm.DB
}

var _ repository.UserRepository = (*UserMysqlRepo)(nil)

func (r *UserMysqlRepo) CreateUser(userInfo *entity.UserInfo) error {
	if userInfo == nil {
		return errors.New("userInfo input is null")
	}
	ctx := context.Background()
	userInfo.UserId = idgenerator.NativeNextID()
	slog.Info(fmt.Sprintf("生成的用户ID: %d", userInfo.UserId))
	db.AutoMigrate(&entity.UserInfo{})
	return gorm.G[entity.UserInfo](db).Create(ctx, userInfo)
}
func (r *UserMysqlRepo) GetUserInfoByUserId(userId int64) (*entity.UserInfo, error) {
	ctx := context.Background()
	userInfo, err := gorm.G[entity.UserInfo](db).Where("user_id = ?", userId).First(ctx)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func NewUserMysqlRepo() repository.UserRepository {
	return &UserMysqlRepo{
		db: db,
	}
}
