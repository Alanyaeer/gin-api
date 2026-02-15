package converter
import (
	"github.com/mitchellh/mapstructure"
	"chat-system/internal/model/dto"
	"chat-system/internal/model/entity"
)

func UserInfoDtoToEntity(userInfoDto *dto.UserInfoReq) (userInfoEntity *entity.UserInfo, err error) {
	var userInfoEntityTemp entity.UserInfo
	if err = mapstructure.Decode(userInfoDto, &userInfoEntityTemp); err != nil {
		return nil, err
	}
	return &userInfoEntityTemp, nil
}