package entity

type UserInfo struct {
	UserId string `json:"userId" gorm:"user_id"`
	Name   string `json:"name" gorm:"name"`
	Age    int    `json:"age" gorm:"age"`
	Sex    string `json:"sex" gorm:"sex"`
}
